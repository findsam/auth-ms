package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string        `bson:"email" json:"email" validate:"required,email"`
	Password string        `bson:"password" json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
	Username string        `bson:"username" json:"username" validate:"required,min=4,max=18,alphanum"`
}

type UserSignInRequest struct {
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}

type UserPublic struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
}

func (u *User) ToPublic() *UserPublic {
	return &UserPublic{
		ID:       u.ID.Hex(),
		Username: u.Username,
		Email:    u.Email,
	}
}