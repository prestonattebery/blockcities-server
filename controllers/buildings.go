package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devinroche/blockcities-server/db"
	"github.com/devinroche/blockcities-server/models"
	"github.com/gorilla/mux"
)

// GetBuildings gets all buildings
func GetBuildings(w http.ResponseWriter, r *http.Request) {
	var buildings []models.Building

	if err := db.DB.Find(&buildings).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&buildings)
}

// GetBuilding gets a single building by id
func GetBuilding(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var building models.Building

	if err := db.DB.First(&building, params["id"]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&building)
}

// CreateBuilding creates a new building
func CreateBuilding(w http.ResponseWriter, r *http.Request) {
	var building models.Building
	json.NewDecoder(r.Body).Decode(&building)

	if err := db.DB.Create(&building); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Panic(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&building)
}
