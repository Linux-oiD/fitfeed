package postgres

import (
	"database/sql"
	"fitfeed/dbm/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Migrator gorm.Migrator

func ConnectToDatabase() (*sql.DB, error) {

	conf := config.Load()
	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
		conf.DB.Postgres.Username,
		conf.DB.Postgres.Password,
		conf.DB.Postgres.DBname,
		conf.DB.Postgres.Host,
		conf.DB.Postgres.Port)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err == nil {
		Migrator = db.Migrator()
		sqlDB, err := db.DB()
		return sqlDB, err
	}
	return nil, err
}
