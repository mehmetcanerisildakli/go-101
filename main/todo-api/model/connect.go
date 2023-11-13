package model

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@(tcp:localhost:6379)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connection is successful")
	return db
}
