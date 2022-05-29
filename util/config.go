package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	ESNodes      []string `mapstructure:"ES_NODES"`
	ESUsername   string   `mapstructure:"ES_USERNAME"`
	ESPassword   string   `mapstructure:"ES_PASSWORD"`
	ESCacertpath string   `mapstructure:"ES_CACERT_PATH"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
