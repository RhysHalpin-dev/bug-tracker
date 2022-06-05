package routes

import (
	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/cmd/controller"
	"github.com/gorilla/mux"
)

func LoginRouteHandler(r *mux.Router) {
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
}

func ProtectedRouteHandler(r *mux.Router) {
	s := r.PathPrefix("/auth/").Subrouter()
	s.Use(controller.AuthMiddleware)
	s.HandleFunc("/profile", controller.ProfileHandler).Methods("POST")
}
