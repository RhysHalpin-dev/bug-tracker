package main

import (
	"net/http"

	"github.com/RhysHalpin-dev/bug-tracker-api/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	//Init Router

	r := mux.NewRouter()
	// init headers
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "Delete"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//Init Routes Welocome
	r.HandleFunc("/", controller.GetWelcome).Methods("GET")
	s := r.PathPrefix("/api/auth").Subrouter()
	s.HandleFunc("/Login", controller.LoginHandler).Methods("POST")
	s.Use(mux.CORSMethodMiddleware(s))

	// Init Listener
	http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(r))
}
