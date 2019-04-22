package dao

import (
	"log"

	. "github.com.br/MarcosPrintes/go_restapi/models"
	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

//struct para conexão com db
type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

//collection que mapearemos com o repositório
const (
	COLLECTION = "movies"
)

//sessão com o db, com métio Dial do mgo, em caso de falha, encerra a sessão, sucesso: passa o valor de session para "db"
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//GetAll deve retornar um json
func (m *MoviesDAO) GetAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesDAO) GetById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MoviesDAO) Create(movie Movie) error {
	err := db.C(COLLECTION).Insert(movie)
	return err
}

func (m *MoviesDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *MoviesDAO) Update(id string, movie Movie) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &movie)
	return err
}
