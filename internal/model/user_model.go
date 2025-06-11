package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string        `bson:"email" json:"email" validate:"required,email"`
	Password string        `json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
}
