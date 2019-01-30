package schedule

import (
	"time"

	"github.com/robfig/cron"
	"gitlab.wallstcn.com/matrix/xgbkb/g"
	"gitlab.wallstcn.com/matrix/xgbkb/std"
)

// A cron expression represents a set of times, using 6 space-separated fields.

// Field name   | Mandatory? | Allowed values  | Allowed special characters
// ----------   | ---------- | --------------  | --------------------------
// Seconds      | Yes        | 0-59            | * / , -
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

var jobsRunner = cron.NewWithLocation(g.ShanghaiTimezone)

func StartJobs() {

	StartJobsTest()

	// 每天更新公司到图数据库
	jobsRunner.AddJob("0 0 1 * * *", std.NewMutexTask(SyncCompaniesAndStocksFromJuyuan).
		WithMutex(std.NewSimpleRedisMutex("SyncCompaniesAndStocksFromJuyuan", time.Minute*4, g.RedisClientMain)))

}

// in local env, we might want to run some tasks upon program start
func StartJobsTest() {
	go func() {
		// SyncCompaniesAndStocksFromJuyuan()
	}()
}

func StopJobs() {
	jobsRunner.Stop()
}
