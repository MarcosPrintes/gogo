package model

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func CheckUser(db *sql.DB) ([]User, error) {
	query := fmt.Sprintf("SELECT * FROM users")
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (user *User) CreateUser(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO users (id, name, password) VALUES ('%d', '%s', '%s')", user.Id, user.Name, user.Password)

	_, err := db.Query(query)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.Id)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
