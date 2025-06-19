package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Email    string        `bson:"email" json:"email" validate:"required,email"`
	Password string        `bson:"password" json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
	Username string        `bson:"username" json:"username" validate:"required,min=3,max=18,alphanum"`
	Meta     *Meta         `bson:"meta" json:"meta"`
	Security *Security     `bson:"security" json:"security"`
}

func (u *User) ToPublic() *UserPublic {
	return &UserPublic{
		ID:       u.ID.Hex(),
		Username: u.Username,
	}
}

func (u *User) ToDatabase() {
	u.Security = NewSecurity()
	u.Meta = NewMeta()
}

type UserSignInRequest struct {
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `bson:"password" json:"password" validate:"required"`
}

type UserPublic struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username" json:"username"`
}
