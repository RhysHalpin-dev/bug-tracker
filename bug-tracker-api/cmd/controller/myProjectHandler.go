package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func MyProjectHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex("61f2c6d504ab8222eb27d92a")
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}}
	//lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "members_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "members"}}}}
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "projects"}, {Key: "localField", Value: "projects"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "returned"}}}}
	showLoadedCursor, err := userCollection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		fmt.Println(err)
	}

	var showsLoaded []bson.D
	if err = showLoadedCursor.All(context.TODO(), &showsLoaded); err != nil {
		panic(err)
	}
	fmt.Println(showsLoaded)
	//test := showsLoaded["email"].(string)
	//Loaded, err := json.Marshal(&jsonLoaded)
	jsonLoaded, err := json.Marshal(&showsLoaded)

	fmt.Println(jsonLoaded)
	if err != nil {
		panic(err)

	} else {
		str := string(jsonLoaded)
		//fmt.Println("stringified json is: ", str)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//JwtRes := model.JwtRes{Message: "Auth successful", Status: 200, Token: jwtToken}
		json.NewEncoder(w).Encode(str)
	}
	//fmt.Println(Loaded)
}
