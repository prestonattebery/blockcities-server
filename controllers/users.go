package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devinroche/blockcities-server/db"
	"github.com/devinroche/blockcities-server/models"
	"github.com/gorilla/mux"
)

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := db.DB.Find(&users).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

// GetUser gets a user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var buildingsArr []models.BuildingInfo
	var bu models.BuildingInfo
	var profile models.Profile
	var u models.User

	rows, err := db.DB.Raw(`
		SELECT building_id, title, address
		FROM user_buildings 
		INNER JOIN buildings ON user_buildings.building_id = buildings.id 
		INNER JOIN users on users.id = user_buildings.user_id WHERE user_id = ?;
		`, params["id"]).Rows()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		db.DB.ScanRows(rows, &bu)
		buildingsArr = append(buildingsArr, bu)
	}

	if err := db.DB.First(&u, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	profile.Name = u.Name
	profile.Username = u.Username
	profile.Owned = buildingsArr

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&profile)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	if err := db.DB.Set("gorm:association_autoupdate", false).Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
