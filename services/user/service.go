package user

import (
	"context"
	"go-api/models"
	"go-api/services/user/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Service interface {
	GetUser(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	CreateUser(ctx context.Context, firstName, lastName, mobile, password, email string ) (*models.User, error)
}

type service struct {
	repository repository.Repository
}
func NewService(repository repository.Repository) Service  {
	return service{repository}
}

func (s service) CreateUser(ctx context.Context, firstName, lastName, mobile, password, email string) (*models.User, error) {
	user := models.User{
		ID: primitive.NewObjectID(),
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		Mobile: mobile,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	return &user, nil
}

func (s service) GetUser(ctx context.Context,id primitive.ObjectID) (*models.User, error)  {
	return s.repository.GetUser(ctx, id)
}
