package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/model"
)

func GetWelcome(w http.ResponseWriter, r *http.Request) {
	// set Header
	welcome := model.Welcome{Message: "Welcome TO Bug Tracker API", Author: "Rhys Halpin", Github: "https://github.com/RhysHalpin-dev"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(welcome)

}
