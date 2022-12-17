package conf

import (
	"github.com/spf13/viper"
	"time"
)

const (
	TIMEOUT_IN_MINUTE = time.Minute
)

type Configuration struct{}

type GlobalConfiguration struct {
	EndPoint EndPoint `mapstructure:"endpoint"`
}

type EndPoint struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func LoadGlobalConfig() (*GlobalConfiguration, error) {
	var gconf GlobalConfiguration
	if err := viper.Unmarshal(&gconf); err != nil {
		return nil, err
	}
	return &gconf, nil
}
