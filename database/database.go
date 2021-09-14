package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Add(fileName string) {
	db, err := sql.Open("mysql", "http@tcp(localhost:3306)/new_ekazan")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	if err != nil {
		log.Fatal(err)
	}

	query := "INSERT INTO webp (name) VALUES(?)"
	res, err := db.Exec(query, fileName)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}
