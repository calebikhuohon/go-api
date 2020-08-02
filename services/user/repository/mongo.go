package repository

import (
	"context"
	"go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Client) Repository  {
	return mongoRepository{
		collection: db.Database("go").Collection("users"),
	}
}

func (m mongoRepository) CreateUser(ctx context.Context, user models.User) error  {
	if _, err := m.collection.InsertOne(ctx, user); err != nil {
		return err
	}
	return  nil
}

func (m mongoRepository) GetUser(ctx context.Context,id primitive.ObjectID) (*models.User, error)  {
	user := new(models.User)

	if err := m.collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(user); err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, errUserNotFound
			}
			return nil, err
	}

	return user, nil
}