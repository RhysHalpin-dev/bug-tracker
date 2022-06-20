package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//IMPORTANT
var userCollection *mongo.Collection
var projectCollection *mongo.Collection

//connect with mongoDB

func init() {
	EnvErr := godotenv.Load("./config/.env")

	if EnvErr != nil {
		fmt.Println("could not load .env file")
		os.Exit(1)
	}

	dbName := os.Getenv("DBNAME")
	userCol := os.Getenv("USERCOL")
	//projectCol := os.Getenv("PROJECTCOL")
	connString := os.Getenv("CONNSTRING")

	//create client options
	clientOptions := options.Client().ApplyURI(connString)

	//connect to mongoDB // pass conext on connection // usingg client option settings
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// if error, fatally log and exit app
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected to mongoDB atlas")

	// if ping successful query databasee names using empty Map{} M: An unordered map. It is the same as D, except it does not preserve order
	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[current Databases Retrieved]:")
	//print out all databases retrieved from ping
	fmt.Println(databases)

	userCollection = client.Database(dbName).Collection(userCol)
	projectCollection = client.Database(dbName).Collection("projects")

	// query collection names using empty Map{} M: An unordered map. It is the same as D, except it does not preserve order
	collections, err := client.Database(dbName).ListCollectionNames(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[current Collections Retrieved]:")
	//print out all collections retrieved from ping
	fmt.Println(collections)
	//collecction instance
	//fmt.Println("collection instance is ready")

	//disconnect from mongoDB database
	/*err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")*/
}
