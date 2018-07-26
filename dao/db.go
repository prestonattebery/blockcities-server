package dao

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/devinroche/blockcities-server/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BlockCitiesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "blockcities"
)

func (m *BlockCitiesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *BlockCitiesDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

func (m *BlockCitiesDAO) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *BlockCitiesDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (m *BlockCitiesDAO) Delete(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

func (m *BlockCitiesDAO) Update(user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}

func RespondWithErr(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
