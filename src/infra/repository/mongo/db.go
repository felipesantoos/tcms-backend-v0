package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type Connector struct{}

func NewConnector() *Connector {
	return &Connector{}
}

var (
	Database          = "tcms_mongo_db"
	ProjectCollection = "project"
)

func (connector Connector) GetConnection() (*mongo.Client, error) {
	uri, err := getMongoConnectionURI()
	if err != nil {
		return nil, err
	}

	var mongoOnce sync.Once
	var clientInstance *mongo.Client
	var clientInstanceError error

	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
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

func getMongoConnectionURI() (string, error) {
	user := "mongo"
	password := "tcmsM0ng0"
	host := "localhost"
	port := "27017"

	connectionUri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", user, password, host, port)

	return connectionUri, nil
}
