package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/devinroche/blockcities-server/config"
	. "github.com/devinroche/blockcities-server/controllers"
	. "github.com/devinroche/blockcities-server/dao"
)

var config = Config{}
var dao = BlockCitiesDAO{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", NewUser).Methods("POST")
	r.HandleFunc("/users", UpdateUser).Methods("PUT")
	r.HandleFunc("/users", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
