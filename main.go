package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Auth-Key", "X-Auth-Secret", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	r := mux.NewRouter()

	r.HandleFunc("/users/{u_id}/building/{b_id}", OwnBuilding).Methods("POST")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")

	r.HandleFunc("/buildings/{id}", GetBuilding).Methods("GET")
	r.HandleFunc("/buildings", GetBuildings).Methods("GET")
	r.HandleFunc("/buildings", CreateBuilding).Methods("POST")

	if err := http.ListenAndServe("localhost:8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
