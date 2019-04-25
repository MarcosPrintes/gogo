package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("go mysql tutorial loco")
	// Open up our database connection.
	// The database is called mytest
	db, err := sql.Open("mysql", "root:13231989@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO mytest VALUES ( 100, 'Marcos Paulo Alho Printes' )")

	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
