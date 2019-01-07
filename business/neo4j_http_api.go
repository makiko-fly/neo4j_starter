package business

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.wallstcn.com/baoer/matrix/xgbkb/g"
	"gitlab.wallstcn.com/baoer/xgbbackend/std/redislogger"
)

// single query
func QueryNeo4j(statement string, params map[string]interface{}) ([]byte, error) {
	parametersStr := "{}"
	if len(params) > 0 {
		bytes, err := json.Marshal(params)
		if err != nil {
			redislogger.Errorf("QueryNeo4j, fails to marshal params: %v, err: %v", params, err)
			return nil, err
		}
		parametersStr = string(bytes)
	}
	reqBodyStr := fmt.Sprintf(singleStatementTempl, statement, parametersStr)
	return callNeo4jHttpApi("/db/data/transaction/commit", reqBodyStr)
}

var singleStatementTempl = `
{
	"statements" : [ {
	  "statement" : "%s",
	  "parameters" : %s
	} ]
  }
`

func callNeo4jHttpApi(path, bodyStr string) ([]byte, error) {
	urlStr := fmt.Sprintf("http://%s:%d%s", g.SysConf.Neo4jDb.Addr, g.SysConf.Neo4jDb.HttpPort, path)
	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer([]byte(bodyStr)))
	if err != nil {
		redislogger.Errorf("callNeo4jHttpApi, create request fails, err: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic bmVvNGo6aWFtc29ycnk=")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		redislogger.Errorf("callNeo4jHttpApi, http POST fails, err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		errStr := fmt.Sprintf("callNeo4jHttpApi, http POST returns code other than 200, code: %d", resp.StatusCode)
		redislogger.Errorf(errStr)
		return nil, errors.New(errStr)
	}
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errStr := fmt.Sprintf("callNeo4jHttpApi, ioutil.ReadAll fails, err: %v", err)
		redislogger.Errorf(errStr)
		return nil, errors.New(errStr)
	}
	return respBodyBytes, nil
}
