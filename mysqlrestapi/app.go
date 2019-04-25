package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// This structure provides references to the router and the database
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Create a database connection and wire up the routes,
func (app *App) Initialize(user, password, dbName string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbName)
	var err error

	app.DB, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// Run app
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

// retrive a user by id, user model.GetUser method
func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		responWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u := user{ID: id}

	if err := u.GetUser(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			responWithError(w, http.StatusNotFound, "user not found") //status 404
		default:
			responWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		responWithJson(w, http.StatusOK, u)
	}

	fmt.Println("getUser => ", u)
}

func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}

	if start < 0 {
		start = 0
	}

	users, err := GetUsers(app.DB, start, count)
	if err != nil {
		responWithError(w, http.StatusInternalServerError, err.Error())
	}

	responWithJson(w, http.StatusOK, users)
}

func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("app createUser")
	fmt.Println("app request => ", r)
	var u user
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&u); err != nil {
		responWithError(w, http.StatusInternalServerError, "create error => "+err.Error())
		return
	}

	defer r.Body.Close()

	if err := u.createUser(app.DB); err != nil {
		responWithError(w, http.StatusBadRequest, err.Error())
	}

	responWithJson(w, http.StatusOK, u)
}

func (app *App) updateUser() {
	fmt.Println("app updateUser")
}

func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("app deleteuser")
}

//initialize api routes
func (app *App) initializeRoutes() {
	fmt.Println("app initialize routes")
	app.Router.HandleFunc("/users", app.getUsers).Methods("GET")
	app.Router.HandleFunc("/newuser", app.createUser).Methods("GET")

}

//===== success/error handlers ===========
func responWithError(w http.ResponseWriter, code int, message string) {
	responWithJson(w, code, map[string]string{"error": message})
}

func responWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
