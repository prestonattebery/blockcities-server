package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/devinroche/blockcities-server/controllers"
	"github.com/devinroche/blockcities-server/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")

	r.HandleFunc("/buildings/{id}", GetBuilding).Methods("GET")
	r.HandleFunc("/buildings", GetBuildings).Methods("GET")
	r.HandleFunc("/buildings", CreateBuilding).Methods("POST")

	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
