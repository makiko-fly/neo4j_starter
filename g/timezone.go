package g

import (
	"time"
)

var ShanghaiTimezone *time.Location

func InitShanghaiTimeZone() {
	tmpTz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	ShanghaiTimezone = tmpTz
}
