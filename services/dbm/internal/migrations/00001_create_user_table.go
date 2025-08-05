package migrations

import (
	"context"
	"database/sql"
	"fitfeed/dbm/internal/models"
	"fitfeed/dbm/internal/postgres"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUserTable, downCreateUserTable)
}

func upCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	if err := postgres.Migrator.CreateTable(&models.Profile{}); err != nil {
		return err
	}
	if err := postgres.Migrator.CreateTable(&models.User{}); err != nil {
		return err
	}
	if err := postgres.Migrator.CreateConstraint(&models.User{}, "Profile"); err != nil {
		return err
	}
	return nil
}

func downCreateUserTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	if err := postgres.Migrator.DropConstraint(&models.User{}, "Profile"); err != nil {
		return err
	}
	if err := postgres.Migrator.DropTable(&models.User{}); err != nil {
		return err
	}
	if err := postgres.Migrator.DropTable(&models.Profile{}); err != nil {
		return err
	}
	return nil
}
