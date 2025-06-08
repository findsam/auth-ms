package model

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Email    string `bson:"email" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,containsany=!@#$%^&*"`
}
