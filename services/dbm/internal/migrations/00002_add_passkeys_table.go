package migrations

import (
	"context"
	"database/sql"
	"fitfeed/dbm/internal/models"
	"fitfeed/dbm/internal/postgres"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddPasskeysTable, downAddPasskeysTable)
}

func upAddPasskeysTable(ctx context.Context, tx *sql.Tx) error {
	if err := postgres.Migrator.CreateTable(&models.Passkey{}); err != nil {
		return err
	}
	if err := postgres.Migrator.CreateConstraint(&models.User{}, "Passkeys"); err != nil {
		return err
	}
	return nil
}

func downAddPasskeysTable(ctx context.Context, tx *sql.Tx) error {
	if err := postgres.Migrator.DropConstraint(&models.User{}, "Passkeys"); err != nil {
		return err
	}
	if err := postgres.Migrator.DropTable(&models.Passkey{}); err != nil {
		return err
	}
	return nil
}
