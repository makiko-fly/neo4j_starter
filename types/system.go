package types

import "fmt"

type SysConfig struct {
	RedisLogger RedisConfig   `yaml:"redis_logger"`
	Neo4jDb     Neo4jDbConfig `yaml:"neo4j_db"`
	Http        HttpConfig    `yaml:"http"`
}

func (self SysConfig) String() string {
	return fmt.Sprintf("==== RedisLogger ====\n%v==== Neo4j ====\n%v", self.RedisLogger, self.Neo4jDb)
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
