package std

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

type EnumEnv int

const (
	EnumEnv_INVALID EnumEnv = iota
	EnumEnv_LOCAL
	EnumEnv_DEV
	EnumEnv_STAGE
	EnumEnv_PROD
)

var Env_LOCAL = "xgblocal"
var Env_DEV = "xgbtest"
var Env_STAGE = "xgbstage"
var Env_PROD = "xgbprod"

func parseEnv(envStr string) EnumEnv {
	if envStr == Env_LOCAL {
		return EnumEnv_LOCAL
	} else if envStr == Env_DEV {
		return EnumEnv_DEV
	} else if envStr == Env_STAGE {
		return EnumEnv_STAGE
	} else if envStr == Env_PROD {
		return EnumEnv_PROD
	}
	return EnumEnv_INVALID
}

var CurEnv EnumEnv

func IsProdEnv() bool {
	return CurEnv == EnumEnv_PROD
}

func LoadConf(config interface{}, path string) error {
	envStr := os.Getenv("CONFIGOR_ENV")
	if envStr == "" {
		envStr = Env_LOCAL
		os.Setenv("CONFIGOR_ENV", envStr)
	}
	env := parseEnv(envStr)
	if env == EnumEnv_INVALID {
		panic(fmt.Sprintf("Invalid env str: %s, env: %v", envStr, env))
	}
	CurEnv = env

	return configor.Load(config, path)
}
