package mongodb

import (
	"context"
	"exchange-rate/config"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func CreateClient() (*mongo.Client, error) {
	config, err := config.GetConfig("config/config")

	if err != nil {
		clientInstanceError = err
	}

	connectionString := fmt.Sprintf(
		config.Mongodb.Database)

	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(connectionString)
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

	return clientInstance, clientInstanceError
}
