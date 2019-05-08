package main

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"

	"log"
	"net/http"

	. "github.com.br/MarcosPrintes/echoApi/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	Echo     *echo.Echo
	DB       *sql.DB
	HashPass *Hash
}

type Hash struct{}

func (h *Hash) encryptPassword(str string) (string, error) {

	saltedBytes := []byte(str)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("create hash error: ", err.Error)
	}

	hashed := string(hashedBytes)

	return hashed, err
}

func (h *Hash) comparePassowrd(hashedStr, str string) error {
	incoming := []byte(str)
	existing := []byte(hashedStr)
	return bcrypt.CompareHashAndPassword(existing, incoming)
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
	app.Echo.POST("/create", app.insert)
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
	var statusCode int
	var res string
	credentials := new(Credentials)
	if err := c.Bind(credentials); err != nil {
		log.Fatal("login bind error", err.Error())
	}
	// row := app.DB.QueryRow("SELECT EXISTS (SELECT * FROM users WHERE user_name = ? AND password = ? )", credentials.UserName, credentials.UserPass)
	rows, err := app.DB.Query("SELECT * FROM users WHERE user_name = ?", credentials.UserName)
	var u User
	for rows.Next() {
		err = rows.Scan(&u.UserId, &u.UserName, &u.UserEmail, &u.UserPass, &u.UserType)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	errPass := app.HashPass.comparePassowrd(u.UserPass, credentials.UserPass)
	if errPass != nil {
		statusCode = http.StatusUnauthorized
		res = "não cadastrado"
	} else {
		statusCode = http.StatusOK
		res = "cadastrado"
	}

	return c.JSON(statusCode, res)
}

func (app *App) insert(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		log.Fatal("bind error", err.Error())
	}

	pass, cryptErr := app.HashPass.encryptPassword(user.UserPass)
	if cryptErr != nil {
		log.Fatal("user crypt error password: ", cryptErr.Error())
	}
	user.UserPass = pass
	fmt.Println("user password", user.UserPass)

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

	sql := fmt.Sprintf("UPDATE users SET user_name = '%s' WHERE user_id = ?", user.UserName)

	stmt, err := app.DB.Prepare(sql)

	if err != nil {
		log.Fatal("update error: ", err.Error())
	}

	res, err := stmt.Exec(user_id)
	if err != nil {
		log.Fatal("update exec error: ", err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
