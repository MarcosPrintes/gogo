package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Db     *sql.DB
}

//=========== START ============
func (app *App) Initialize(user, password, dbName string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbName)
	var err error
	app.Db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.InitializeRoutes()
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/u/{id}", app.getUser)
}

//============= HANDLERS ==============
func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	fmt.Println(param)
	id, err := strconv.Atoi(param["id"])

	if err != nil {
		respondWithJsonError(w, http.StatusBadRequest, err.Error())
	}

	u := User{Id: id}

	if err != nil {
		respondWithJsonError(w, http.StatusInternalServerError, err.Error())
	}

	if err := u.getUser(app.Db); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithJsonError(w, http.StatusNotFound, "user not found") //status 404
		default:
			respondWithJsonError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJson(w, http.StatusOK, u)
}

func fastHttpHandler(ctx *fasthttp.RequestCtx) {

}

//======== JSON RESPONSES =======
func respondWithJsonError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, message)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
