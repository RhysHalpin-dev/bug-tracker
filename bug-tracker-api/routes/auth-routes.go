package routes

import (
	"github.com/RhysHalpin-dev/bug-tracker-api/controller"
	"github.com/gorilla/mux"
)

func AuthRouteHandler(r *mux.Router) {
	r.HandleFunc("/Login", controller.LoginHandler).Methods("POST")
}
