package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBname   string
}

func ConnectToDatabase(conf PGConfig) (*gorm.DB, error) {

	dbURL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
		conf.Username,
		conf.Password,
		conf.DBname,
		conf.Host,
		conf.Port)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	return db, err
}
