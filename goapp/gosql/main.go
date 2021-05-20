package main

// http://go-database-sql.org/index.html

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VividCortex/mysqlerr"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"root:test@tcp(mysql:3306)/training_DB") //DB_HOST : mysql (not localhost)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO items(name, price, created_at, updated_at) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Dolly", 120, "2013-10-01", "2013-10-01") //Queryとの違いに注意！
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	// 一般的に、他のプログラミング言語や外部ライブラリでは、SQLで空の結果となってもエラーや例外が発生しないことが多いが、 .QueryRow() では空の結果を考慮する必要がある
	var name string
	err = db.QueryRow("select name from items where id = ?", 10).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No row selected")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)

	_, err_to_checknum := db.Query("SELECT someval FROM sometable")
	if driverErr, ok := err_to_checknum.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {

		}
	}
}
