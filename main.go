package main

import (
	"github.com/api-projet/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	// USER
	router.HandleFunc("/api/user/", controllers.FetchUser).Methods("GET")
	router.HandleFunc("/api/user/new/", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login/", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/delete/{uuid}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/user/put/{uuid}", controllers.PutUser).Methods("PUT")

	// VOTE
	router.HandleFunc("/api/vote/", controllers.FetchVote).Methods("GET")
	router.HandleFunc("/api/vote/create/", controllers.CreateVote).Methods("POST")
	router.HandleFunc("/api/vote/update/{uuid}", controllers.EditVote).Methods("PUT")
	router.HandleFunc("/api/vote/show/{uuid}", controllers.SingleVote).Methods("GET")
	router.HandleFunc("/api/vote/delete/{uuid}", controllers.DeleteVote).Methods("DELETE")
	router.HandleFunc("/api/vote/submit/", controllers.SubmitVote).Methods("PUT")


	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8081" //localhost
	}

	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)) //Launch the app, visit localhost:8888/api

	if err != nil {
		fmt.Print(err)
	}
}
