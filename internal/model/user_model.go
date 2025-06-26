package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Socials struct {
	Twitter  string `bson:"twitter" json:"twitter"`
	LinkedIn string `bson:"linkedin" json:"linkedin"`
	GitHub   string `bson:"github" json:"github"`
}

type UserDetails struct {
	Description string   `bson:"description" json:"description"`
	Location    string   `bson:"location" json:"location"`
	Socials     *Socials `bson:"socials" json:"socials"`
}

type User struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string        `bson:"email" json:"email" validate:"required,email"`
	Password string        `bson:"password" json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
	Username string        `bson:"username" json:"username" validate:"required,min=3,max=18,alphanum"`
	Socials  *Socials      `bson:"socials" json:"socials"`
	Meta     *Meta         `bson:"meta" json:"meta"`
	Security *Security     `bson:"security" json:"security"`
}

func (u *User) ToPublic() *UserPublic {
	return &UserPublic{
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
	Id       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string        `bson:"username" json:"username"`
}
