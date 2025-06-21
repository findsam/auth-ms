package repo

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const PAYMENT_DB_NAME = "payment"

type PaymentRepository interface{}

type PaymentRepositoryImpl struct {
	db *mongo.Database
}

func NewPaymentRepositoryImpl(db *mongo.Database) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}
