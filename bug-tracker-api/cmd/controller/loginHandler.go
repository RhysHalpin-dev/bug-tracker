package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/model"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Login
	var mongoUser bson.M
	//var convertedResult model.Login

	// parse and decode request body into Login struct // throw error if not possible
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil { // if JSON is not correctly formated BAD REQUEST
		fmt.Println("Error: ", err)
		status := model.Status{Message: "Bad Request", Status: 400}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(status)
	}

	if err == nil {
		//retrieve document matching the users email
		filter := bson.M{"email": user.Email}
		err := collection.FindOne(context.TODO(), filter).Decode(&mongoUser)
		fmt.Println("Found user document: ", mongoUser)
		// log error if retrival is unsuccessful
		if err != nil {
			status := model.Status{Message: "Bad Request", Status: 400}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status)
		}

		//convert retrieved mongoDB bson to Login struct
		//bsonBytes, _ := bson.Marshal(result)
		//bson.Unmarshal(bsonBytes, &convertedResult)

		//compare password to hashed password
		match := CheckPasswordHash(user.Password, mongoUser["password"].(string))
		fmt.Println("Password correct? ", match)

		//compare user given password and retrieved result password //demopassword123
		if !match {
			status := model.Status{Message: "Auth unSuccessful", Status: 401}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(status)

		} else { // Clean request no error
			jwtToken := generateToken(mongoUser)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			JwtRes := model.JwtRes{Message: "Auth successful", Status: 200, Token: jwtToken}
			json.NewEncoder(w).Encode(JwtRes)

		}
	}
}
