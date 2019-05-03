package model

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (c *Credentials) Signup(db *sql.DB) (User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE user_name like '%s' AND password like '%s'", c.User, c.Password)
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("signup error => ", err.Error())
	}
	var user User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.UserType); err != nil {
			log.Fatal("signup scan error =>", err.Error())
		}
	}
	return user, err
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	query := fmt.Sprintf("SELECT * FROM users")
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	users := []User{}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.UserType); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (user *User) CreateUser(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO users (user_name, user_email, password, user_type) VALUES ('%s', '%s', '%s', '%s')", user.Name, user.Email, user.Password, user.UserType)

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
