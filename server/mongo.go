package server

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func ConnectToMongo() (client *mongo.Client, err error)  {
	ctx, cancel := context.WithTimeout(context.Background(),  10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return  client, nil
}

func ensureMongoIndexes(db *mongo.Database)   {

}
