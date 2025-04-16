package infrastructure

import (
	"message/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


type MongoClient struct {
	Db *mongo.Database
}

func NewMongoClient(mongoConfig config.MongoConfig) (*MongoClient, error) {
	opts_auth := options.Credential{
		Username: mongoConfig.Username,
		Password: mongoConfig.Password,
	}
	client_opts := options.Client().SetAuth(opts_auth).SetHosts([]string{mongoConfig.Host})
	client, err := mongo.Connect(client_opts)
	if err != nil {
		return nil, err
	}
	return &MongoClient{Db: client.Database(mongoConfig.Database)}, nil
}
