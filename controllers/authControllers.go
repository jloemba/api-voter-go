package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/api-projet/models"
	u "github.com/api-projet/utils"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	
	fmt.Println(account)
	
	//fmt.Println(err)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

var PutUser = func(w http.ResponseWriter, r *http.Request) {
/*
	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	if len(param["uuid"]) < 1 {
		fmt.Println("L'url n'a pas de uuid")
		return
	}

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	key := param["uuid"]

	resp := account.PutUserHandler(key)
	u.Respond(w, resp)
*/


	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur

	params := mux.Vars(r)
	fmt.Println(reflect.TypeOf(params))
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.PutUserHandler(params["id"],account) //Create account
	u.Respond(w, resp)
}


var DeleteUser = func(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	if len(param["uuid"]) < 1 {
		fmt.Println("L'url n'a pas de uuid")
		return
	}
	key := param["uuid"]

	resp := models.DeleteUserHandler(key)
	u.Respond(w, resp)
}


