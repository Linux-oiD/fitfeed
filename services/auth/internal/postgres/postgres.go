package postgres

import (
	"fitfeed/auth/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(conf *config.AppConfig) (*gorm.DB, error) {

	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
		conf.DB.Postgres.Username,
		conf.DB.Postgres.Password,
		conf.DB.Postgres.DBname,
		conf.DB.Postgres.Host,
		conf.DB.Postgres.Port)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	return db, err
}
