package migrations

import (
	"context"
	"database/sql"
	"fitfeed/dbm/internal/db"
	"fitfeed/dbm/internal/models"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateInitialTables, downCreateInitialTables)
}

func upCreateInitialTables(ctx context.Context, tx *sql.Tx) error {
	driver := ctx.Value("driver").(string)
	gdb, err := db.GetGormTx(tx, driver)
	if err != nil {
		return err
	}

	return gdb.Migrator().CreateTable(
		&models.User{},
		&models.Profile{},
		&models.OauthProvider{},
	)
}

func downCreateInitialTables(ctx context.Context, tx *sql.Tx) error {
	driver := ctx.Value("driver").(string)
	gdb, _ := db.GetGormTx(tx, driver)
	return gdb.Migrator().DropTable(models.OauthProvider{}, &models.Profile{}, &models.User{})
}
