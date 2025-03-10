package mongo_client

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

)


type MongoClient struct {
	db *mongo.Database
}

func NewMongoClient(host string) (*MongoClient, error) {
	opts_auth := options.Credential{
		Username: "username",
		Password: "password",
	}
	client_opts := options.Client().SetAuth(opts_auth).SetHosts([]string{host})
	client, err := mongo.Connect(client_opts)
	if err != nil {
		return nil, err
	}
	return &MongoClient{db: client.Database("beep")}, nil
}
