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

	var buildingsArr []models.BuildingInfo
	var bu models.BuildingInfo
	var profile models.Profile

	rows, _ := db.DB.Raw(`
		SELECT building_id, title, address
		FROM user_buildings 
		INNER JOIN buildings ON user_buildings.building_id = buildings.id 
		INNER JOIN users on users.id = user_buildings.user_id WHERE user_id = ?;
		`, params["id"]).Rows()

	defer rows.Close()

	for rows.Next() {
		db.DB.ScanRows(rows, &bu)
		buildingsArr = append(buildingsArr, bu)
	}

	var u models.User
	db.DB.First(&u, params["id"])

	profile.Name = u.Name
	profile.Username = u.Username
	profile.Owned = buildingsArr

	json.NewEncoder(w).Encode(&profile)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Set("gorm:association_autoupdate", false).Create(&user)
	json.NewEncoder(w).Encode(&user)
}
