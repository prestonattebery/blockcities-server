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

	// r.HandleFunc("/building/{id}", GetBuilding).Methods("GET")
	// r.HandleFunc("/buildings", GetBuildings).Methods("GET")
	// r.HandleFunc("/buildings", CreateBuilding).Methods("POST")

	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}

// // get all users
// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	var users []User
// 	db.Find(&users)
// 	json.NewEncoder(w).Encode(&users)
// }

// // get a user by id
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var user User
// 	db.First(&user, params["id"])
// 	json.NewEncoder(w).Encode(&user)
// }

// // create new user
// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	fmt.Println(&user)
// 	db.Create(&user)
// 	json.NewEncoder(w).Encode(&user)
// }
