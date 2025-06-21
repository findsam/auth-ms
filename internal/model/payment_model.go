package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)


type CreatePaymentBody struct {
    StoreId string `json:"store_id" validate:"required"`
    OwnerId  string `json:"owner_id" validate:"required"`
    Sub int `json:"amount" validate:"required"` 
}

type Payment struct {
    ID      bson.ObjectID `bson:"_id,omitempty" json:"id"`
    StoreId bson.ObjectID `bson:"store_id" json:"store_id" validate:"required"`
    Amount  float64       `bson:"amount" json:"amount" validate:"required"`
	Meta     *Meta        `bson:"meta" json:"meta"`
}