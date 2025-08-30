package main

import (
	"fitfeed/auth/internal/config"
	httpcontroller "fitfeed/auth/internal/controller/http"
	"fitfeed/auth/internal/repo/oauthdb"
	"fitfeed/auth/internal/repo/profiledb"
	"fitfeed/auth/internal/repo/userdb"
	"fitfeed/auth/internal/usecase/oauthmanager"
	"fitfeed/auth/internal/usecase/profilemanager"
	"fitfeed/auth/internal/usecase/usermanager"
	"fmt"
	"log"
	"net/http"

	"fitfeed/auth/pkg/httpserver"
	"fitfeed/auth/pkg/postgres"
)

func main() {

	conf := config.Load()
	db, err := postgres.ConnectToDatabase(postgres.PGConfig{
		Host:     conf.DB.Postgres.Host,
		Port:     conf.DB.Postgres.Port,
		Username: conf.DB.Postgres.Username,
		Password: conf.DB.Postgres.Password,
		DBname:   conf.DB.Postgres.DBname,
	})
	if err != nil {
		log.Fatal("DB error")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	um := usermanager.New(userdb.New(db))
	om := oauthmanager.New(oauthdb.New(db))
	pm := profilemanager.New(profiledb.New(db))

	srv := httpserver.New(conf.Auth.Port)
	srv.Handler = httpcontroller.New(um, om, pm)

	done := make(chan bool, 1)

	go httpserver.GracefulShutdown(srv, done)

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")

}
