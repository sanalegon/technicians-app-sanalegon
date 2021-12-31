package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://sanalegon:1234@cluster0.tqs48.mongodb.net/technicians-app?retryWrites=true&w=majority")

/* ConnectDB is the function that allow us to connect the DB */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error was found")
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("error when trying to ping db")
		log.Fatal(err.Error())
		return client
	}

	log.Println("Succesful connection to the DB")
	return client
}

/* CheckConnection is a ping to the DB */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return 0
	}

	return 1
}
