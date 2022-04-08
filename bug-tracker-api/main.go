package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RhysHalpin-dev/bug-tracker-api/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	EnvErr := godotenv.Load("./config/.env")

	if EnvErr != nil {
		fmt.Println("could not load .env file")
		os.Exit(1)
	}
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
	http.ListenAndServe(os.Getenv("GOPORT"), handlers.CORS(headers, methods, origins)(r))
}
