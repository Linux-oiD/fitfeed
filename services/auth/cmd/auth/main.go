package main

import (
	"fitfeed/auth/internal/config"
	"fitfeed/auth/internal/server"
	"fmt"
	"net/http"
)

func main() {

	conf := config.Load()
	srv := server.Init(conf)
	srvURL := fmt.Sprintf(":%d", conf.Web.Port)

	http.ListenAndServe(srvURL, srv)
}
