package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"fitfeed/dbm/internal/config"
	"fitfeed/dbm/internal/db"
	_ "fitfeed/dbm/internal/migrations"

	"github.com/pressly/goose/v3"
)

func main() {

	conf := config.Load()
	var handler slog.Handler
	if conf.IsProd {
		handler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		handler = slog.NewTextHandler(os.Stdout, nil)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)

	fset := flag.NewFlagSet("dbm", flag.ExitOnError)
	dir := fset.String("dir", ".", "migration directory")
	fset.Parse(os.Args[1:])

	args := fset.Args()
	if len(args) < 1 {
		fset.Usage()
		return
	}

	nativeDB, err := db.Connect(conf)
	if err != nil {
		slog.Error("failed to connect to db", "error", err)
		os.Exit(1)
	}
	defer nativeDB.Close()

	if err := goose.SetDialect(conf.DB.Driver); err != nil {
		slog.Error("failed to set dialect", "error", err)
		os.Exit(1)
	}

	ctx := context.WithValue(context.Background(), "driver", conf.DB.Driver)
	command := args[0]

	slog.Info("running migration command", "command", command, "driver", conf.DB.Driver)

	if err := goose.RunContext(ctx, command, nativeDB, *dir, args[1:]...); err != nil {
		slog.Error("goose command failed", "command", command, "error", err)
		os.Exit(1)
	}
	slog.Info("migration completed successfully")
}
