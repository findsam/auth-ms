package repo

import (
	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const PAYMENT_DB_NAME = "payment"

type PaymentRepository interface{
	Create() (*model.Payment, error)
}

type PaymentRepositoryImpl struct {
	db *mongo.Database
}

func NewPaymentRepositoryImpl(db *mongo.Database) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (r *PaymentRepositoryImpl) Create() (*model.Payment, error) {
	return nil,nil
}
