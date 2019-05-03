package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com.br/MarcosPrintes/plc/model"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
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

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	app.InitializeRoutes()

	http.ListenAndServe(":8089", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(app.Router))

}

func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/api/v1/all", app.getAllUsers).Methods("GET")
	app.Router.HandleFunc("/api/v1/signup", app.signup).Methods("POST")
	app.Router.HandleFunc("/api/v1/createUser", app.createUser).Methods("POST")
}

func (app *App) signup(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	decode := json.NewDecoder(r.Body)

	if err := decode.Decode(&credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonResponseError(w, http.StatusBadRequest, err.Error())
	}

	defer r.Body.Close()

	user, err := credentials.Signup(app.Db)
	if err != nil {
		log.Fatal("deu erro ai hein!")
		jsonResponseError(w, http.StatusInternalServerError, err.Error())
	}
	if user.Name == "" {
		jsonResponseSuccess(w, http.StatusUnauthorized, "Usuário não encontrado")
	}
	jsonResponseSuccess(w, http.StatusOK, user)
}

func (app *App) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := GetAllUsers(app.Db)
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

//=============== Responses =================
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
