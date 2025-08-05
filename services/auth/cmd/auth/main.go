package main

import (
	"fitfeed/auth/internal/config"
	"fitfeed/auth/internal/database"
	"fmt"
)

func main() {
	conf := config.Load()
	fmt.Printf("conf: %+v\n", conf)
	database.Init()

}
