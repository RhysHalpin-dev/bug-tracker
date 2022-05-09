package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/RhysHalpin-dev/bug-tracker-api/model"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

/*
	Generate a JWT Token uisng user data
*/
func generateToken(mongoUser bson.M) (tokenString string) {

	hmacSecretKey := []byte(os.Getenv("SECRETKEY"))
	//create new token and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"client_id": mongoUser["_id"].(primitive.ObjectID),
		"email":     mongoUser["email"].(string),
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecretKey)
	fmt.Println(tokenString, err)
	AuthorizationJwt(tokenString)
	return tokenString

}

func AuthorizationJwt(jwtToken string) {
	hmacSecretKey := []byte(os.Getenv("SECRETKEY"))
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["userName"], claims["admin"])
	} else {
		fmt.Println(err)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	var UserObject model.UserObject
	var mongoUser bson.M

	// parse and decode request body into Login struct // throw error if not possible
	fmt.Println("body", r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&UserObject)
	fmt.Println(UserObject)

	if err != nil {

		//retrieve document matching the users email
		objectID, _ := primitive.ObjectIDFromHex(UserObject.UserObject)

		filter := bson.M{"_id": objectID}
		err := collection.FindOne(context.TODO(), filter).Decode(&mongoUser)

		println(mongoUser)

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
