package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/RhysHalpin-dev/bug-tracker/bug-tracker-api/model"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

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
	hmacSecretKey := []byte(os.Getenv("SECRETKEY")) //jwt secret signing key stored in env

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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// take jwt token from request
		reqToken := r.Header.Get("Authorization")
		jwtToken := strings.Split(reqToken, "Bearer")
		// check token is split in 2
		if len(jwtToken) != 2 {
			status := model.Status{Message: "Bad Request", Status: 400}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(status)
		} else {
			//store split token value after bearer
			reqToken = strings.TrimSpace(jwtToken[1])
			fmt.Println(reqToken)

			hmacSecretKey := []byte(os.Getenv("SECRETKEY")) //jwt secret signing key stored in env

			token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return hmacSecretKey, nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				//generate valid action response
				fmt.Println(claims["client_id"], claims["email"])
				//call next handler
				next.ServeHTTP(w, r)

			} else {
				// generate user not valid responce
				fmt.Println(err)
				status := model.Status{Message: "Bad Request", Status: 400}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(status)
				return
			}
		}
	})

}
