package user

import "go-api/models"

type getUserResponse struct {
	User *models.User `json:"user"`
}

type createUserResponse struct {
	User *models.User `json:"user"`
}
