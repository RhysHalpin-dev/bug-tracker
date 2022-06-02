package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	var UserObject model.UserObject
	var mongoUser bson.M

	// parse and decode request body into Login struct // throw error if not possible
	fmt.Println("body", r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UserObject)
	//fmt.Println(UserObject)

	if err == nil {

		//retrieve document matching the users email
		objectID, _ := primitive.ObjectIDFromHex(UserObject.UserObject)

		filter := bson.M{"_id": objectID}
		err := collection.FindOne(context.TODO(), filter).Decode(&mongoUser)

		fmt.Println("Found user document: ", mongoUser)

		if err != nil {
			fmt.Println("Error: ", err)
			status := model.Status{Message: "Bad Request", Status: 400}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			Profile := model.Profile{Email: mongoUser["email"].(string), Name: mongoUser["name"].(string), Bio: mongoUser["bio"].(string)}
			json.NewEncoder(w).Encode(Profile)
		}
	}
}
