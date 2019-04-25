package main

func main() {

	var a = App{}

	a.Initialize("root", "", "rest_api_example")

	a.Run(":8080")

	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/rest_api_example")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer db.Close()

	// // insert, err := db.Query("INSERT INTO users VALUES ( 200, 'Marcos', 30 )")
	// del, err := db.Query("DELETE FROM users")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// // defer insert.Close()
	// defer del.Close()

}
