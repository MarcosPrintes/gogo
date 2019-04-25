package main

import (
	"database/sql"
	"fmt"

	. "github.com.br/MarcosPrintes/gomysql/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("go mysql tutorial loco")
	// Open up our database connection.
	// The database is called mytest
	// db, err := sql.Open("mysql", "root:13231989@tcp(127.0.0.1:3306)/test")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO mytest VALUES ( 200, 'Mel' )")

	// del, err := db.Query("DELETE FROM mytest")

	results, err := db.Query("SELECT * FROM mytest")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User
		err := results.Scan(&user.ID, &user.NAME)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("User => ", user)
	}

	//be careful deferring Queries if you are using transactions
	defer results.Close()
}
