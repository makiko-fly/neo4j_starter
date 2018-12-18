package std

import (
	"os"

	"github.com/jinzhu/configor"
)

func LoadConf(config interface{}, path string) error {
	if os.Getenv("CONFIGOR_ENV") == "" {
		os.Setenv("CONFIGOR_ENV", "local")
	}
	return configor.Load(config, path)
}
