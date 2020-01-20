package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/api-projet/models"
	u "github.com/api-projet/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var CreateVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}
	err := json.NewDecoder(r.Body).Decode(vote) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := vote.Create()
	u.Respond(w, resp)
}


var EditVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}
	
	err := json.NewDecoder(r.Body).Decode(vote) //decode the request body into struct and failed if any error occur
	
	params := mux.Vars(r)["uuid"]
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	
	resp := models.UpdateVote(params,vote)
	u.Respond(w, resp)
}


var DeleteVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}
		
	params := mux.Vars(r)["uuid"]


	resp := models.DeleteVote(params,vote)
	u.Respond(w, resp)
}

var SingleVote = func(w http.ResponseWriter, r *http.Request) {

	vote := &models.Vote{}


	params := mux.Vars(r)["uuid"]

	
	resp := models.SingleVote(params,vote)
	u.Respond(w, resp)
}

var SubmitVote = func(w http.ResponseWriter, req *http.Request) {

	vote := &models.Vote{}
	err := json.NewDecoder(req.Body).Decode(vote)
	fmt.Println(vote)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	
	resp := models.SubmitVote(vote.UUIDVote,vote)
	u.Respond(w, resp)
}

var FetchVote = func(w http.ResponseWriter, r *http.Request){
	resp := models.FetchVote()
	u.Respond(w, resp)
}

