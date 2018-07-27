package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	. "github.com/devinroche/blockcities-server/dao"
	. "github.com/devinroche/blockcities-server/models"
	"github.com/gorilla/mux"
)

var dao = BlockCitiesDAO{}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := dao.FindById(params["id"])
	if err != nil {
		RespondWithErr(w, http.StatusInternalServerError, "Invalid ID")
		return
	}
	RespondWithJSON(w, http.StatusOK, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dao.FindAll()
	if err != nil {
		RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, users)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithErr(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	user.ID = bson.NewObjectId()
	if err := dao.Insert(user); err != nil {
		RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithErr(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	if err := dao.Update(user); err != nil {
		RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithErr(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	if err := dao.Delete(user); err != nil {
		RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
