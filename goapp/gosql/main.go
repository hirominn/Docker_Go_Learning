package main

// http://go-database-sql.org/index.html

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"root:test@tcp(mysql:3306)/training_DB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
