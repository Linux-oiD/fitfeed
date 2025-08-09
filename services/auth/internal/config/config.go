package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Auth struct {
		Port   int    `mapstructure:"port"`
		Secret string `mapstructure:"secret"`
	} `mapstructure:"auth"`

	DB struct {
		Driver string `mapstructure:"driver"`

		Postgres struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			DBname   string `mapstructure:"dbname"`
		} `mapstructure:"postgres"`
	} `mapstructure:"database"`

	Web struct {
		Hostname string `mapstructure:"hostname"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"web"`
}

func Load() *AppConfig {

	viper.SetConfigName("auth-config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../../config") // Look for the config file in the module root directory
	viper.AddConfigPath("./config")     // Look for the config file in the current directory

	viper.AutomaticEnv()
	viper.SetEnvPrefix("fitfeed")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &config
}
