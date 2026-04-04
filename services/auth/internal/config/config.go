package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	IsProd bool `mapstructure:"is_prod"`
	Auth   struct {
		Port      int    `mapstructure:"port"`
		Prefix    string `mapstructure:"prefix"`
		Secret    string `mapstructure:"secret"`
		MaxAge    int    `mapstructure:"max_session_age"`
		Providers map[string]struct {
			Enabled      bool   `mapstructure:"enabled"`
			ClientID     string `mapstructure:"client_id"`
			ClientSecret string `mapstructure:"client_secret"`
		} `mapstructure:"providers"`
	} `mapstructure:"auth"`
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
	Web struct {
		Hostname string `mapstructure:"hostname"`
		Protocol string `mapstructure:"protocol"`
		Port     int    `mapstructure:"port"`
	} `mapstructure:"web"`
}

func Load() *AppConfig {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	// 1. Look for config file in path provided by env FITFEED_CONF
	if envConf := os.Getenv("FITFEED_CONF"); envConf != "" {
		viper.AddConfigPath(envConf)
	}

	// 2. Look in the current working directory
	viper.AddConfigPath(".")

	// 3. Fallback for local development relative to the service root
	viper.AddConfigPath("../../")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("fitfeed")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &config
}
