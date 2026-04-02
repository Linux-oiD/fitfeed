package main

import (
	"fitfeed/auth/internal/config"
	httpcontroller "fitfeed/auth/internal/controller/http"
	"fitfeed/auth/internal/oauth"
	"fitfeed/auth/internal/repo/oauthdb"
	"fitfeed/auth/internal/repo/profiledb"
	"fitfeed/auth/internal/repo/userdb"
	"fitfeed/auth/internal/usecase/jwtmanager"
	"fitfeed/auth/internal/usecase/oauthmanager"
	"fitfeed/auth/internal/usecase/profilemanager"
	"fitfeed/auth/internal/usecase/usermanager"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"fitfeed/auth/pkg/httpserver"
	"fitfeed/auth/pkg/postgres"
)

func main() {

	conf := config.Load()

	// Initialize slog
	var handler slog.Handler
	if conf.Auth.IsProd {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, nil)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)

	oauth.NewAuth(conf)
	db, err := postgres.ConnectToDatabase(postgres.PGConfig{
		Host:     conf.DB.Postgres.Host,
		Port:     conf.DB.Postgres.Port,
		Username: conf.DB.Postgres.Username,
		Password: conf.DB.Postgres.Password,
		DBname:   conf.DB.Postgres.DBname,
	})
	if err != nil {
		logger.Error("DB connection error", "error", err)
		os.Exit(1)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	udb := userdb.New(db)
	pdb := profiledb.New(db)
	odb := oauthdb.New(db)

	um := usermanager.New(udb, pdb, logger)
	om := oauthmanager.New(odb, logger)
	pm := profilemanager.New(pdb, logger)
	jm := jwtmanager.New(conf.Auth.Secret, time.Duration(conf.Auth.MaxAge)*time.Second)

	srv := httpserver.New(conf.Auth.Port)
	srv.Handler = httpcontroller.New(um, om, pm, jm)

	done := make(chan bool, 1)

	go httpserver.GracefulShutdown(srv, done)

	logger.Info("Starting server...", "port", conf.Auth.Port)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("http server error", "error", err)
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	logger.Info("Graceful shutdown complete.")

}
