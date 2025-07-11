package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Store struct {
	Id      bson.ObjectID `bson:"_id,omitempty" json:"id"`
	OwnerId bson.ObjectID `bson:"owner_id,omitempty" json:"owner_id"`
	Tiers   *[]Tier       `json:"tiers" bson:"tiers"`
	Meta    *Meta         `json:"meta" bson:"meta"`
}

type Tier struct {
	Amount      float64  `json:"amount" bson:"amount"`
	Description string   `json:"description" bson:"description"`
	Benefits    []string `json:"benefits" bson:"benefits"`
}

type StorePublic struct {
	Id 	bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Tiers *[]Tier `json:"tiers" bson:"tiers"`
}

type UserStoreResult struct {
	User  UserPublic     `json:"user" bson:"user"`
	Store StorePublic 	`json:"store" bson:"store"`
}

func (s *Store) ToPublic() StorePublic {
	return StorePublic{
		Id:    s.Id,
		Tiers: s.Tiers,
	}
}
