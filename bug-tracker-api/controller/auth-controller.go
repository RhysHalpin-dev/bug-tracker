package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RhysHalpin-dev/bug-tracker-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetWelcome(w http.ResponseWriter, r *http.Request) {
	// set Header
	welcome := model.Welcome{Message: "Welcome TO Bug Tracker API", Author: "Rhys Halpin", Github: "https://github.com/RhysHalpin-dev"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(welcome)

}

/*
 *Compare the users input password the
 * hashed version stored on DB
 *
 */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Login
	var result bson.M
	var convertedResult model.Login

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
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		fmt.Println("Found user document: ", result)
		// log error if retrival is unsuccessful
		if err != nil {
			status := model.Status{Message: "Bad Request", Status: 400}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status)
		}

		//convert retrieved mongoDB bson to Login struct
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &convertedResult)

		//compare password to hashed password
		match := CheckPasswordHash(user.Password, convertedResult.Password)
		fmt.Println("Password correct? ", match)

		//compare user given password and retrieved result password //demopassword123
		if !match {
			status := model.Status{Message: "Auth unSuccessful", Status: 401}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(status)

		} else { // Clean request no error

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			status := model.Status{Message: "Auth successful", Status: 200}
			json.NewEncoder(w).Encode(status)

		}
	}
}
