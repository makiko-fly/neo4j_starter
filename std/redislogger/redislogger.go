package redislogger

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/redis.v5"
)

const KeyLogList string = "kb:custom:log:list"
const MaxListSize = 20000

var redisClient *redis.Client

// type RedisLoggerConf struct {
// 	Host        string
// 	Port        string
// 	Password    string
// 	DB          int
// 	MaxIdle     int
// 	IdleTimeout int
// }

func Init(host string, port int64, password string) error {
	address := fmt.Sprintf("%s:%d", host, port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})

	if err := redisClient.Ping().Err(); err != nil {
		fmt.Printf("==> Failed to ping redis server %v, err: %v\n", address, err)
		return err
	} else {
		fmt.Println("==> Successfully initialized redis logger")
		Printf("Redislogger initialized at %s", time.Now().Format("2006-01-02 15:04:05"))
	}

	return nil
}

func getKey(slot int32) string {
	if slot <= 0 {
		return KeyLogList
	}
	return fmt.Sprintf("%s:%d", KeyLogList, slot)
}

func WrapLogLine(trackId int32, line string) string {
	var hostName string
	if tmpStr, err := os.Hostname(); err != nil {

	} else {
		// get last 5 characters of host
		if len(tmpStr) > 5 {
			hostName = tmpStr[len(tmpStr)-5:]
		} else {
			hostName = tmpStr
		}
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	var finalStr string
	if trackId > 0 {
		finalStr = fmt.Sprintf("[%s] [%s] [%d] %s", timeStr, hostName, trackId, line)
	} else {
		finalStr = fmt.Sprintf("[%s] [%s] %s", timeStr, hostName, line)
	}
	return finalStr
}

func logToRedis(slotNum int32, wrappedLine string) error {
	if newSize, err := redisClient.LPush(getKey(slotNum), wrappedLine).Result(); err != nil {
		return err
	} else if newSize > MaxListSize {
		redisClient.RPop(getKey(slotNum))
	}
	return nil
}

func Printf(format string, v ...interface{}) {
	var slotNum int32 = 0
	var trackId int32 = 0
	line := fmt.Sprintf(format, v...)
	wrappedLine := WrapLogLine(trackId, line)
	logToRedis(slotNum, wrappedLine)
	// print normal message to console
	fmt.Println(wrappedLine)
}

func Errf(format string, v ...interface{}) {
	var slotNum int32 = 0
	var trackId int32 = 0
	line := fmt.Sprintf(format, v...)
	wrappedLine := WrapLogLine(trackId, line)
	logToRedis(slotNum, wrappedLine)
	// print error message to console, error channel
	fmt.Fprintln(os.Stderr, wrappedLine)
}

func Slotf(slotNum int32, format string, v ...interface{}) {
	var trackId int32 = 0
	line := fmt.Sprintf(format, v...)
	wrappedLine := WrapLogLine(trackId, line)
	logToRedis(slotNum, wrappedLine)
	// print to console
	fmt.Println(wrappedLine)
}

func SlotErrf(slotNum int32, format string, v ...interface{}) {
	var trackId int32 = 0
	line := fmt.Sprintf(format, v...)
	wrappedLine := WrapLogLine(trackId, line)
	logToRedis(slotNum, wrappedLine)
	// print error message to console, error channel
	fmt.Fprintln(os.Stderr, wrappedLine)
}

func Entries(slotNum, limit int32) ([]string, error) {
	if limit <= 0 {
		limit = -1
	}
	if entries, err := redisClient.LRange(getKey(slotNum), 0, int64(limit)).Result(); err != nil {
		return nil, err
	} else {
		return entries, nil
	}
}

// =====================================================================================================================
// http API

func MountRedisLog(g *echo.Group) {
	redislogGroup := g.Group("/redislog")
	redislogGroup.GET("", redisLogHandler)
}

func redisLogHandler(c echo.Context) error {
	slot, _ := strconv.ParseInt(c.FormValue("slot"), 10, 64)
	limit, _ := strconv.ParseInt(c.FormValue("limit"), 10, 64)
	if limit <= 0 {
		limit = 5000
	}
	entries, err := Entries(int32(slot), int32(limit))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	} else {
		if compiledHtmlTempate == nil {
			if tmpTemplate, err := template.New("redislog").Parse(HTML_TEMPLATE); err != nil {
				return err
			} else {
				compiledHtmlTempate = tmpTemplate
			}
		}
		var dest bytes.Buffer
		if err := compiledHtmlTempate.Execute(&dest, entries); err != nil {
			return err
		}
		return c.HTML(http.StatusOK, dest.String())
	}
}

var HTML_TEMPLATE = `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>服务器端错误日志</title>
    <style type="text/css">
    	p {
    		margin: 0;
    	}
    </style>
</head>

<body>
<div class='main-520-wrapper'>
	<div class="xgb_content">
			{{ range . }}
			<p>
            {{ . }}
			</p>
			{{ end }}
	</div>
</div>
</body>
`
var compiledHtmlTempate *template.Template
