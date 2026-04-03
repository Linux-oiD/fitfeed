package v1

import (
	"encoding/json"
	"fitfeed/api/internal/config"
	"net/http"
)

type ConfigController struct {
	conf *config.AppConfig
}

func NewConfigController(conf *config.AppConfig) *ConfigController {
	return &ConfigController{conf: conf}
}

type WebConfig struct {
	AuthURL string `json:"auth_url"`
	APIURL  string `json:"api_url"`
}

func (c *ConfigController) GetConfig(w http.ResponseWriter, r *http.Request) {
	webConf := WebConfig{
		AuthURL: "http://localhost:8081", // In prod get from conf
		APIURL:  "http://localhost:8082", // In prod get from conf
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(webConf)
}
