package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"imgoptimizer/config"
)

type Db struct {
	con *sql.DB
}

func (db *Db) Close() {
	if db != nil && db.con != nil {
		db.con.Close()
	}
}

func (db *Db) Add(fileName string) (int64, error) {
	query := "INSERT IGNORE INTO webp (name) VALUES (?)"
	res, err := db.con.Exec(query, fileName)

	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func New(conf *config.Conf) (*Db, error) {
	var db Db
	var err error

	if conf.Database.Connection == "" {
		return nil, fmt.Errorf("нет данных для подключения")
	}

	db.con, err = sql.Open("mysql", conf.Database.Connection)
	if err != nil {
		return nil, err
	}

	if err = db.con.Ping(); err != nil {
		return nil, err
	}

	return &db, nil
}
