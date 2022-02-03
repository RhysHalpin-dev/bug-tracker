package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RhysHalpin-dev/bug-tracker-api/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetWelcome(w http.ResponseWriter, r *http.Request) {
	// set Header
	welcome := model.Welcome{Message: "Welcome TO Bug Tracker API", Author: "Rhys Halpin", Github: "https://github.com/RhysHalpin-dev"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(welcome)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Login
	var result bson.M
	var result1 model.Login

	decoder := json.NewDecoder(r.Body)
	// parse and decode request body into Login struct // throw error if not possible
	err := decoder.Decode(&user)

	if err != nil { // if JSON is not correctly formated BAD REQUEST
		status := model.Status{Message: "Bad Request", Status: 400}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status)
	}

	if err == nil {
		//retrieve document matching the users email
		filter := bson.M{"email": user.Email}
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		fmt.Println(result)
		// log error if retrival is unsuccessful
		if err != nil {
			status := model.Status{Message: "Bad Request", Status: 400}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status)
		}

		//convert retrieved mongoDB bson to Login struct
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &result1)

		//compare user given password and retrieved result password
		if user.Password != result1.Password {
			status := model.Status{Message: "Auth unSuccessful", Status: 404}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status)

		} else { // Clean request no error

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			status := model.Status{Message: "Auth successful", Status: 200}
			json.NewEncoder(w).Encode(status)

		}
	}
}
