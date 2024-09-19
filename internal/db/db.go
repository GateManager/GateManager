package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func NewMySqlStorage(driver string, cfg mysql.Config) (*sql.DB, error) {
	var db *sql.DB
	var err error

	if driver == "sqlite3" {
		db, err = sql.Open("sqlite3", "./GateManager.db")
		if err != nil {
			return nil, err
		}
	} else {
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
