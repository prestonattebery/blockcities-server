package main

import (
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8080" //localhost
	}

	r.HandleFunc("/users/{u_id}/building/{b_id}", OwnBuilding).Methods("POST")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/signin", LoginUser).Methods("POST")

	r.HandleFunc("/buildings/{id}", GetBuilding).Methods("GET")
	r.HandleFunc("/buildings", GetBuildings).Methods("GET")
	r.HandleFunc("/buildings", CreateBuilding).Methods("POST")

	if err := http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
