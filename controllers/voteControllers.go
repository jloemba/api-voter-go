package controllers

import (
	"fmt"
	"reflect"
	"encoding/json"
	"github.com/api-projet/models"
	u "github.com/api-projet/utils"	
	"github.com/gorilla/mux"
	"net/http"
)

var CreateVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}
	
	err := json.NewDecoder(r.Body).Decode(vote) //decode the request body into struct and failed if any error occur
	
	
	//fmt.Println(err)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := vote.Create() //Create account
	u.Respond(w, resp)
}


var EditVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}
	
	err := json.NewDecoder(r.Body).Decode(vote) //decode the request body into struct and failed if any error occur
	
	params := mux.Vars(r)["uuid"]
	fmt.Println(reflect.TypeOf(params))
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	
	resp := models.UpdateVote(params,vote) //Create account
	u.Respond(w, resp)
}


var DeleteVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}
		
	params := mux.Vars(r)["uuid"]

	fmt.Println(reflect.TypeOf(params))

	resp := models.DeleteVote(params,vote) //Create account
	u.Respond(w, resp)
}



var SingleVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}
		
	params := mux.Vars(r)["uuid"]

	resp := models.SingleVote(params,vote) //Create account
	u.Respond(w, resp)
}


var SubmitVote = func(w http.ResponseWriter, r *http.Request) {

	//fmt.Println("deeeeeee")

	//fmt.Println(mux.Vars(r))

	uuidvote := mux.Vars(r)["uuidvote"]
	uuidaccount := mux.Vars(r)["uuidaccount"]
	token := mux.Vars(r)["token-user"]

	//fmt.Println("deeeeeee 2")
	resp := models.SubmitVote(uuidvote,uuidaccount,token) //Create account
	//fmt.Println("deeeeeee 3")

	u.Respond(w, resp)
}