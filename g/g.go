package g

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gitlab.wallstcn.com/matrix/xgbkb/std/redislogger"
	"gitlab.wallstcn.com/matrix/xgbkb/types"
	"gopkg.in/redis.v5"
)

var SysConf types.SysConfig

var RedisClientMain *redis.Client

var JuyuanDb *gorm.DB

func InitRedisClients() {
	addr := fmt.Sprintf("%s:%d", SysConf.RedisMain.Host, SysConf.RedisMain.Port)
	RedisClientMain = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: SysConf.RedisMain.Auth,
		DB:       int(SysConf.RedisMain.DB),
	})
	if err := RedisClientMain.Ping().Err(); err != nil {
		panic(err)
	}
}

func InitDb() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		SysConf.MysqlJuyuan.User,
		SysConf.MysqlJuyuan.Password,
		SysConf.MysqlJuyuan.Host,
		SysConf.MysqlJuyuan.Port,
		SysConf.MysqlJuyuan.DbName,
	)

	tmpDb, err := gorm.Open("mysql", connStr)
	if err != nil {
		redislogger.Errf("Failed to open mysql connection, err: %v", err)
		panic(err)
	} else {
		redislogger.Printf("Successfully connect to db with conf: %v", SysConf.MysqlJuyuan)
		JuyuanDb = tmpDb
		// set log mode
		JuyuanDb.LogMode(SysConf.MysqlJuyuan.LogMode)
		JuyuanDb.DB().SetMaxIdleConns(int(SysConf.MysqlJuyuan.MaxIdle))
		JuyuanDb.DB().SetMaxOpenConns(int(SysConf.MysqlJuyuan.MaxConn))
	}
}
