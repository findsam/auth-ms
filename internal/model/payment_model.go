package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreatePaymentBody struct {
	StoreId string `json:"store_id" validate:"required"`
	Tier    int    `json:"tier" validate:"gte=0,lte=3"`
}

type Payment struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	StoreId  bson.ObjectID `bson:"store_id" json:"store_id" validate:"required"`
	StripeId string        `bson:"stripe_id" json:"stripe_id" validate:"required"`
	Meta     *Meta         `bson:"meta" json:"meta"`
}

type PaymentAggregateResult struct {
	Payment *Payment    `json:"payment" bson:"payment"`
	// User    *UserPublic `json:"user" bson:"user"`
	Store   *Store      `json:"store" bson:"store"`
}
