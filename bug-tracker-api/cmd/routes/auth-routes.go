package routes

import (
	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/cmd/controller"
	"github.com/gorilla/mux"
)

func AuthRouteHandler(r *mux.Router) {
	r.HandleFunc("/auth/login", controller.LoginHandler).Methods("POST")
}

func ProfileRouteHandler(r *mux.Router) {
	r.HandleFunc("/auth/profile", controller.ProfileHandler).Methods("POST")
}
