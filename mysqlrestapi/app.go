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
	// log.Fatal(http.ListenAndServe(addr, app.Router))
	http.ListenAndServe(addr, app.Router)
}

//initialize api routes
func (app *App) initializeRoutes() {
	fmt.Println("app initialize routes")
	app.Router.HandleFunc("/all", app.getUsers)
	app.Router.HandleFunc("/user/{id}", app.getUser)
	app.Router.HandleFunc("/create", app.createUser)
	app.Router.HandleFunc("/del/{id}", app.deleteUser)
	app.Router.HandleFunc("/upd/{id}", app.updateUser)
}

// retrive a user by id, user model.GetUser method
func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u := user{ID: id}

	if err := u.GetUser(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "user not found") //status 404
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJson(w, http.StatusOK, u)
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
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJson(w, http.StatusOK, users)
}

// This handler assumes that the request body is a JSON object containing the details of the user to be created
func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("app createUser")
	var u user
	// decoder := json.NewDecoder(r.Body)

	// if err := decoder.Decode(&u); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "create error, invalid request payload => "+err.Error())
	// 	return
	// }

	defer r.Body.Close()

	if err := u.createUser(app.DB); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJson(w, http.StatusCreated, u)
}

func (app *App) updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("app updateUser")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "update error: "+err.Error())
		return
	}

	var u user
	decoder := json.NewDecoder(r.Body)
	fmt.Println("app updateUser decoder:", decoder)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "update decoder erro:"+err.Error())
		return
	}
	defer r.Body.Close()
	u.ID = id
	if err := u.updateUser(app.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, u)
}

func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "delete error: "+err.Error())
		return
	}

	u := user{ID: id}

	if err := u.deleteUser(app.DB); err != nil {
		respondWithError(w, http.StatusBadRequest, "delete error:"+err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"result:": "success"})
}

//===== success/error handlers ===========
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}
