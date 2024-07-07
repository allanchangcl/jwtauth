package database

import (
	"database/sql"
	"fmt"
	"log"
)

func Connection() {
	db, err := sql.Open("sqlite3", "./jwtauth.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connected")
	defer db.Close()
}
