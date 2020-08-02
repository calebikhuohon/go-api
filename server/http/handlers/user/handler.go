package user

import (
	"context"
	"github.com/gorilla/mux"
	"go-api/errors"
	"go-api/server/http/decoder"
	"go-api/server/http/responses"
	"go-api/services/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) error
	GetUser(w http.ResponseWriter, r *http.Request) error
}

type handler struct {
	userService user.Service
}

type createUserRequestBody struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Mobile string `json:"mobile"`
	Password string `json:"password"`
	Email string `json:"email"`
}
func NewUserHandler(userService user.Service) Handler  {
	return handler{userService}
}

func (u handler) CreateUser(w http.ResponseWriter, r *http.Request) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	requestBody := &createUserRequestBody{}
	if err := decoder.DecodeJSON(r.Body, requestBody); err != nil {
		return err
	}

	newUser, err := u.userService.CreateUser(
		ctx,
		requestBody.FirstName,
		requestBody.LastName,
		requestBody.Mobile,
		requestBody.Password,
		requestBody.Email,
		)

	if err != nil {
		return err
	}

	return responses.OK("User registration successful", createUserResponse{User: newUser}).ToJSON(w)
}

func (u handler) GetUser(w http.ResponseWriter, r *http.Request) error  {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	id := mux.Vars(r)["userId"]
	objectId, idErr := primitive.ObjectIDFromHex(id)

	if idErr != nil {
		return errors.Error("Invalid user ID")
	}
	user, err := u.userService.GetUser(ctx, objectId)

	if err != nil {
		return err
	}

	return responses.OK("User retrieval successful", getUserResponse{user}).ToJSON(w)
}


