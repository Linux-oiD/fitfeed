package database

import (
	"fitfeed/auth/internal/config"
	"log"
)

func Init() {
	conf := config.Load()
	log.Printf("config: %+v\n", conf)
}
