package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DB struct {
		Driver   string `mapstructure:"driver"`
		Postgres struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			DBname   string `mapstructure:"dbname"`
		} `mapstructure:"postgres"`
	} `mapstructure:"database"`
}

func Load() AppConfig {
	// Set up viper to read the config.yaml file
	viper.SetConfigName("dbm-config")   // Config file name without extension
	viper.SetConfigType("toml")         // Config file type
	viper.AddConfigPath("./config")     // Look for the config file in the current directory
	viper.AddConfigPath("../../config") // Look for the config file in the current directory

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

	return config
}
