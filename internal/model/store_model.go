package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Store struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	OwnerId     primitive.ObjectID  `bson:"owner_id,omitempty" json:"owner_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	Tiers       *[]Tier       `json:"tiers" bson:"tiers"`
	Meta        *Meta         `json:"meta" bson:"meta"`
}

type Tier struct {
	Amount      int64    `json:"amount" bson:"amount"`
	Description string   `json:"description" bson:"description"`
	Benefits    []string `json:"benefits" bson:"benefits"`
}
