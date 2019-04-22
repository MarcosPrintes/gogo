package movierouter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	. "github.com.br/MarcosPrintes/go_restapi/config/dao"
	. "github.com.br/MarcosPrintes/go_restapi/models"
)

var dao = MoviesDAO{}

// Método que pode ser chamado no caso de algum erro
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJson(w, code, map[string]string{"error": message})
}

// Método que pode ser chamado no caso de sucesso, retornar em JSON
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all")
	movies, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.GetById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Create: Invalid request payload")
		return
	}

	movie.ID = bson.NewObjectId()
	if err := dao.Create(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var movie Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Upadate: Invalid request payload")
		return
	}

	if err := dao.Update(params["id"], movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respondWithJson(w, http.StatusOK, map[string]string{"update result": movie.ID + "atualizado com sucesso."})
	respondWithJson(w, http.StatusOK, map[string]string{"update result": "atualizado com sucesso."})

}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, map[string]string{"delete result": "delete success."})
}
