package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
