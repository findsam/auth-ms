package model


type User struct {
	ID   	string `bson:"_id,omitempty" json:"id"`
	Email 	string `bson:"email" json:"email"`
}