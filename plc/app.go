package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com.br/MarcosPrintes/plc/model"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Db     *sql.DB
}

func (app *App) Initialize(user, password, dbName string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbName)
	fmt.Println(connectionString)
	var err error
	app.Db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	app.Router = mux.NewRouter()
	app.InitializeRoutes()
}

func (app *App) Run(addr string) {
	http.ListenAndServe(addr, app.Router)
}

func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/all", app.checkUser).Methods("GET")
	app.Router.HandleFunc("/createUser", app.createUser).Methods("POST")
}

func (app *App) checkUser(w http.ResponseWriter, r *http.Request) {
	users, err := CheckUser(app.Db)
	if err != nil {
		jsonResponseError(w, http.StatusInternalServerError, err.Error())
	}
	jsonResponseSuccess(w, http.StatusOK, users)
}

func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	decode := json.NewDecoder(r.Body)
	// create user by request body
	if err := decode.Decode(&user); err != nil {
		jsonResponseError(w, http.StatusInternalServerError, "invalid request payload => "+err.Error())
	}
	defer r.Body.Close()

	if err := user.CreateUser(app.Db); err != nil {
		jsonResponseError(w, http.StatusBadRequest, "cannot create users, error => "+err.Error())
	}
	jsonResponseSuccess(w, http.StatusOK, user)
}

func jsonResponseError(w http.ResponseWriter, code int, message string) {
	jsonResponseSuccess(w, code, map[string]string{})
}

func jsonResponseSuccess(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
