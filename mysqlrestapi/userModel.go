package main

import (
	"database/sql"
	"fmt"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// methods that will query in db
func (user *user) GetUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", user.ID)
	return db.QueryRow(statement).Scan(&user.Name, &user.Age)
}

func GetUsers(db *sql.DB, start, count int) ([]user, error) {
	statement := fmt.Sprintf("SELECT id, name, age FROM users LIMIT %d OFFSET %d", count, start)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			fmt.Println("next")
		}
		users = append(users, u)
	}

	return users, nil
}

func (user *user) createUser(db *sql.DB) error {
	fmt.Println("model createUser")
	// statement := fmt.Sprintf("INSERT INTO users (name, age) VALUES ('%s','%d')", "Outro Alguem ", 987)
	statement := fmt.Sprintf("INSERT INTO users (name, age) VALUES ('%s','%d')", "Outro Alguem ", 987)

	_, err := db.Query(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (user *user) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", user.ID)
	_, err := db.Exec(statement)
	return err
}

func (user *user) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET name='%s', age='%d' WHERE id='%d'", user.Name, user.Age, user.ID)
	_, err := db.Exec(statement)
	return err
}
