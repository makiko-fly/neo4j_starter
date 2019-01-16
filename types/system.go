package types

import "fmt"

type SysConfig struct {
	RedisLogger RedisConfig   `yaml:"redis_logger"`
	Neo4jDb     Neo4jDbConfig `yaml:"neo4j_db"`
	Http        HttpConfig    `yaml:"http"`
	RedisMain   RedisConfig   `yaml:"redis_main"`
	MysqlJuyuan MysqlConfig   `yaml:"mysql_juyuan"`
}

func (self SysConfig) String() string {
	return fmt.Sprintf("==== RedisLogger ====\n%s==== Neo4j ====\n%s==== Redis Main ====\n%s==== JuyuanMysql ====\n%s",
		self.RedisLogger, self.Neo4jDb, self.RedisMain, self.MysqlJuyuan)
}

type Neo4jDbConfig struct {
	Addr      string
	BoltPort  int64
	HttpPort  int64
	HttpsPort int64
}

func (self Neo4jDbConfig) String() string {
	return fmt.Sprintf(" Address: %s,\n BoltPort: %d,\n HttpPort: %d,\n HttpsPort: %d\n", self.Addr, self.BoltPort,
		self.HttpPort, self.HttpsPort)
}

type RedisConfig struct {
	Host        string
	Port        int64
	Auth        string
	DB          int64
	MaxIdle     int64
	IdleTimeout int64
}

func (self RedisConfig) String() string {
	return fmt.Sprintf(" Host: %s,\n Port: %d,\n DB: %d,\n MaxIdle: %d,\n IdleTimeout: %d\n", self.Host, self.Port,
		self.DB, self.MaxIdle, self.IdleTimeout)
}

type HttpConfig struct {
	Port int64
}

func (self HttpConfig) String() string {
	return fmt.Sprintf("  Port: %d", self.Port)
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	DbName   string `yaml:"db_name"`
	MaxIdle  int64  `yaml:"max_idle"`
	MaxConn  int64  `yaml:"max_conn"`
	LogMode  bool   `yaml:"log_mode"`
}

func (self MysqlConfig) String() string {
	return fmt.Sprintf(" User: %s,\n Host: %s,\n Port: %d,\n DbName: %s,\n LogMode: %v", self.User, self.Host,
		self.Port, self.DbName, self.LogMode)
}
