package main

import (
	"fitfeed/api/internal/config"
	httpcontroller "fitfeed/api/internal/controller/http"
	"fitfeed/api/internal/repo/profiledb"
	"fitfeed/api/internal/repo/userdb"
	"fitfeed/api/internal/usecase/usermanager"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"fitfeed/api/pkg/httpserver"
	"fitfeed/api/pkg/postgres"
)

func main() {

	conf := config.Load()

	// Initialize slog
	var handler slog.Handler
	if conf.IsProd {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, nil)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)

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

	um := usermanager.New(udb, pdb, logger)

	srv := httpserver.New(conf.API.Port)
	srv.Handler = httpcontroller.New(um, conf)

	done := make(chan bool, 1)

	go httpserver.GracefulShutdown(srv, done)

	logger.Info("Starting api server...", "port", conf.API.Port)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error("http server error", "error", err)
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	logger.Info("Graceful shutdown complete.")

}
