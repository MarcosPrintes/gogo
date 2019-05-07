package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"

	. "github.com.br/MarcosPrintes/echoApi/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	Echo *echo.Echo
	DB   *sql.DB
}

func main() {
	var err error
	var app App
	app.DB, err = sql.Open("mysql", "root:@/places") // user:password@/dbname
	if err != nil {
		log.Fatal(err.Error())
	}

	defer app.DB.Close()

	app.Echo = echo.New()

	app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.POST, echo.GET, echo.OPTIONS, echo.PUT, echo.DELETE, echo.HEAD},
	}))

	app.Echo.GET("/get", app.getTest)
	app.Echo.POST("/login", app.login)
	app.Echo.DELETE("/del/:id", app.delTest)
	app.Echo.PUT("/upd/:id", app.updateTest)
	app.Echo.Logger.Fatal(app.Echo.Start(":6543"))
}

func (app *App) getTest(c echo.Context) error {
	rows, err := app.DB.Query("SELECT * FROM users")
	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserPass, &user.UserType)
		if err != nil {
			log.Fatal("query error: ", err.Error())
		}
		users = append(users, user)
	}
	return c.JSON(http.StatusOK, users)
}

func (app *App) login(c echo.Context) error {
	credentials := new(Credentials)
	if err := c.Bind(credentials); err != nil {
		log.Fatal("login bind error", err.Error())
	}

	sql := fmt.Sprintf("SELECT * FROM users WHERE user_name like '%s' AND password like '%s' ", credentials.UserName, credentials.UserPass)

	query, err := app.DB.Query(sql)

	if err != nil {
		log.Fatal("select erro: ", err.Error())
	}
	var user User
	for query.Next() {
		err := query.Scan(&user.UserId, &user.UserName, &user.UserPass, &user.UserEmail, &user.UserType)
		if err != nil {
			log.Fatal("login erro: ", err.Error())
		}
	}
	return c.JSON(http.StatusOK, user)
}

func (app *App) insert(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		log.Fatal("bind error", err.Error())
	}
	res, err := app.DB.Exec("INSERT INTO users (user_name, user_email, password, user_type) VALUES (?, ?, ?, ?)", user.UserName, user.UserEmail, user.UserPass, user.UserType)
	if err != nil {
		log.Fatal("Insert error: ", err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (app *App) delTest(c echo.Context) error {
	user_id := c.Param("id")
	sql := "DELETE FROM users WHERE user_id = ?"
	stmt, err := app.DB.Prepare(sql)
	if err != nil {
		log.Fatal("prepare error: ", err.Error())
	}
	res, err := stmt.Exec(user_id)
	if err != nil {
		log.Fatal("delete error: ", err.Error())
	}
	fmt.Println(res.RowsAffected())
	return c.JSON(http.StatusOK, res)
}

func (app *App) updateTest(c echo.Context) error {
	user_id := c.Param("id")
	user := new(User)
	if err := c.Bind(user); err != nil {
		log.Fatal("update error: ", err.Error())
	}
	fmt.Println("name", user.UserName)
	sql := fmt.Sprintf("UPDATE users SET user_name = '%s' WHERE user_id = ?", user.UserName)
	stmt, err := app.DB.Prepare(sql)
	if err != nil {
		log.Fatal("update error: ", err.Error())
	}

	res, err := stmt.Exec(user_id)
	if err != nil {
		log.Fatal("update exec error: ", err.Error())
	}

	fmt.Println("res => ", res)
	return c.JSON(http.StatusOK, "upd")
}
