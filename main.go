package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/api-projet/app"
	"github.com/api-projet/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts
	//router.HandleFunc("/api/me/accounts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts
	router.HandleFunc("/api/vote/create", controllers.CreateVote).Methods("POST")
	router.HandleFunc("/api/vote/update/{uuid}", controllers.EditVote).Methods("PUT")
	router.HandleFunc("/api/vote/show/{uuid}", controllers.SingleVote).Methods("GET")
	router.HandleFunc("/api/vote/delete/{uuid}", controllers.DeleteVote).Methods("DELETE")

	router.HandleFunc("/api/vote/submit", controllers.SubmitVote).Methods("POST")


	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888" //localhost
	}

	//fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8888/api
	if err != nil {
		//fmt.Print(err)
	}
}
