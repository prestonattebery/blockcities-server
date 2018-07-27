package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/devinroche/blockcities-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=blockcities sslmode=disable password=password")

	if err != nil {
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Building{})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")

	r.HandleFunc("/building/{id}", GetBuilding).Methods("GET")
	r.HandleFunc("/buildings", GetBuildings).Methods("GET")
	r.HandleFunc("/buildings", CreateBuilding).Methods("POST")

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

// get all buildings
func GetBuildings(w http.ResponseWriter, r *http.Request) {
	var buildings []Building
	db.Find(&buildings)
	json.NewEncoder(w).Encode(&buildings)
}

// get a building by id
func GetBuilding(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var building Building
	db.First(&building, params["id"])
	json.NewEncoder(w).Encode(&building)
}

// create a new building
func CreateBuilding(w http.ResponseWriter, r *http.Request) {
	var building Building
	json.NewDecoder(r.Body).Decode(&building)

	db.Create(&building)
	json.NewEncoder(w).Encode(&building)
}
