package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"username"`
	Email string `json:"email"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")

	output, err := json.Marshal(NewUser)

	fmt.Println(string(output))

	if err != nil {
		fmt.Println("deu erro no json")
	}

	query := "INSERT INTO users set user_nickname='" + NewUser.Name + "',	user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "',user_email='" + NewUser.Email + "'"
	sql.OpenDB()
	q, err := database.Exec(query)

	if err != nil {
		fmt.Println("deu erro no sql")
	}

	fmt.Println(query)
}

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/api/user/create", CreateUser).Methods("GET")
}
