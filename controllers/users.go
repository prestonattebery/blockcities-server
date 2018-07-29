package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/devinroche/blockcities-server/db"
	"github.com/devinroche/blockcities-server/models"
	"github.com/gorilla/mux"
)

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// GetUser gets a user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}
