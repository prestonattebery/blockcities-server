package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/devinroche/blockcities-server/db"
	"github.com/devinroche/blockcities-server/models"
	"github.com/gorilla/mux"
)

// GetBuildings gets all buildings
func GetBuildings(w http.ResponseWriter, r *http.Request) {
	var buildings []models.Building
	db.DB.Find(&buildings)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&buildings)
}

// GetBuilding gets a single building by id
func GetBuilding(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var building models.Building
	db.DB.First(&building, params["id"])
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&building)
}

// CreateBuilding creates a new building
func CreateBuilding(w http.ResponseWriter, r *http.Request) {
	var building models.Building
	json.NewDecoder(r.Body).Decode(&building)
	db.DB.Create(&building)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&building)
}
