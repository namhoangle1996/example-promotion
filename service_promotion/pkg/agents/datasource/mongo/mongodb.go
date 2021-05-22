package mongo

import (
	"gitlab.com/wallet-gpay/service-promotion/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBconnection(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}
	err = client.Connect(helpers.ContextWithTimeOut())

	if err != nil {
		panic("can't be reachable server")
	}

	return client
}
