package main

import (
	"net/http"

	"github.com/RhysHalpin-dev/bug-tracker-api/controller"
	"github.com/gorilla/mux"
)

func main() {

	//Init Router

	r := mux.NewRouter()

	//Init Routes Welocome
	r.HandleFunc("/", controller.GetWelcome).Methods("GET")
	s := r.PathPrefix("/api/auth").Subrouter()
	s.HandleFunc("/Login", controller.LoginHandler).Methods("POST")

	// Init Listener
	http.ListenAndServe(":8000", r)
}
