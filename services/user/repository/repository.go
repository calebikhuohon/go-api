package repository

import (
	"context"
	"go-api/errors"
	"go-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	errUserNotFound = errors.NotFound("user")
)

type Repository interface {
	GetUser(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	CreateUser(ctx context.Context, user models.User) error
}
