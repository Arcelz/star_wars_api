package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once
var DB = "star_wars"

const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	PLANETS          = "planets"
)

func connect() {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
}

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	connect()
	return clientInstance, clientInstanceError
}

//GetMongoCollection - Return mongodb collection to work with
func GetMongoCollection(collection string) (*mongo.Collection, error) {
	connect()
	if clientInstanceError != nil {
		return nil, clientInstanceError
	}
	return clientInstance.Database(DB).Collection(collection), nil
}
