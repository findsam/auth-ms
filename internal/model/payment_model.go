package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)


type Payment struct {
    ID      bson.ObjectID `bson:"_id,omitempty" json:"id"`
    StoreId bson.ObjectID `bson:"store_id" json:"store_id" validate:"required"`
    Amount  float64       `bson:"amount" json:"amount" validate:"required"`
}