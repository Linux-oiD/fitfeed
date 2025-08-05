package main

import (
	"context"
	"database/sql"
	"fitfeed/dbm/internal/config"
	_ "fitfeed/dbm/internal/migrations"
	"fitfeed/dbm/internal/postgres"
	"flag"
	"log"
	"os"

	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("dbm", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("dbm: failed to parse flags: %v", err)
	}

	args := flags.Args()
	conf := config.Load()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	var db *sql.DB
	var err error

	if conf.DB.Driver == "postgres" {

		db, err = postgres.ConnectToDatabase()
		if err != nil {
			log.Fatalf("dbm: failed to open DB: %v", err)
		}

		defer func() {
			if err := db.Close(); err != nil {
				log.Fatalf("dbm: failed to close DB: %v", err)
			}
		}()

	} else {
		db = nil
		log.Fatalf("dbm: driver %s is not supported", conf.DB.Driver)
	}

	command := args[0]

	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	ctx := context.Background()
	if err := goose.RunContext(ctx, command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
