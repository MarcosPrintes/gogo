package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (user *User) getUser(db *sql.DB) error {
	strGetUser := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", user.Id)
	return db.QueryRow(strGetUser).Scan(&user.Name, &user.Age)
}
