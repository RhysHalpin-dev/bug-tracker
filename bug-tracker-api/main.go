package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/cmd/controller"
	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/cmd/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	//Load .env befor router init
	EnvErr := godotenv.Load("./config/.env")

	if EnvErr != nil {
		fmt.Println("could not load .env file")
		os.Exit(1)
	}
}
func main() {
	//Init Router

	r := mux.NewRouter()
	// init headers
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "Delete"})
	origins := handlers.AllowedOrigins([]string{"*"})

	//Init Routes Welocome
	r.HandleFunc("/", controller.GetWelcome).Methods("GET")
	//Create Authenticaion and Authorization endpoint
	s := r.PathPrefix("/apiv1/").Subrouter()
	// Pass above subrouter to routes handler
	routes.LoginRouteHandler(s)
	routes.ProtectedRouteHandler(s)
	//r.Use(mux.CORSMethodMiddleware(s))

	// Init Listener
	http.ListenAndServe(os.Getenv("GOPORT"), handlers.CORS(headers, methods, origins)(r))
}
