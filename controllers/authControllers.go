package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/api-projet/models"
	u "github.com/api-projet/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message(false, "requete non valide"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "requete non valide"))
		return
	}
	fmt.Println(account)
	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

var PutUser = func(w http.ResponseWriter, r *http.Request) {


	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur

	params := mux.Vars(r) 
	if err != nil {
		u.Respond(w, u.Message(false, "requete non valide"))
		return
	}

	resp := models.PutUserHandler(params["uuid"],account) //Create account
	u.Respond(w, resp)
}


var DeleteUser = func(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	if len(param["uuid"]) < 1 {
		u.Respond(w, u.Message(false, "l'url n'a pas d'id"))
		return
	}
	key := param["uuid"]

	resp := models.DeleteUserHandler(key)
	u.Respond(w, resp)
}

var FetchUser = func(w http.ResponseWriter, r *http.Request){
	resp := models.FetchUser()
	u.Respond(w, resp)
}
