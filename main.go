package main

import (
	"fmt"
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
	router.HandleFunc("/api/user/delete/{uuid}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/user/put/{uuid}", controllers.PutUser).Methods("PUT")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8888/api
	if err != nil {
		fmt.Print(err)
	}
}
