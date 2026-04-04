package db

import (
	"database/sql"
	"fitfeed/dbm/internal/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(conf *config.AppConfig) (*sql.DB, error) {
	var dialector gorm.Dialector

	switch conf.DB.Driver {
	case "postgres":
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
			conf.DB.Postgres.Username, conf.DB.Postgres.Password,
			conf.DB.Postgres.DBname, conf.DB.Postgres.Host, conf.DB.Postgres.Port)
		dialector = postgres.Open(dsn)
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DB.Mysql.Username, conf.DB.Mysql.Password,
			conf.DB.Mysql.Host, conf.DB.Mysql.Port, conf.DB.Mysql.DBname)
		dialector = mysql.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", conf.DB.Driver)
	}

	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gdb.DB()
}

func GetGormTx(tx *sql.Tx, driver string) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch driver {
	case "postgres":
		dialector = postgres.New(postgres.Config{Conn: tx})
	case "mysql":
		dialector = mysql.New(mysql.Config{Conn: tx})
	default:
		return nil, fmt.Errorf("unsupported driver: %s", driver)
	}
	return gorm.Open(dialector, &gorm.Config{})
}
