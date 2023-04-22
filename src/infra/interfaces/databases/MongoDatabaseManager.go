package databases

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatabaseManager interface {
	GetConnection() (*mongo.Client, error)
}
