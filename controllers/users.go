package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/devinroche/blockcities-server/db"
	"github.com/devinroche/blockcities-server/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

// UpdateUser updates a user given ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User
	fmt.Println(params["id"])
	db.DB.First(&user, params["id"])

	fmt.Println(&user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := db.DB.Set("gorm:association_autoupdate", false).Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

// LoginUser logs user in given username and password
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	var user models.User

	json.NewDecoder(r.Body).Decode(&login)

	err := db.DB.Where("username = ?", login.Username).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Password = ""

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func OwnBuilding(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	db.DB.Exec(`
		INSERT INTO user_buildings (user_id, building_id)
		VALUES (?, ?)`, params["u_id"], params["b_id"])

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w)
}
