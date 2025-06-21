package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const PAYMENT_DB_NAME = "payment"

type PaymentRepository interface{
	Create(sid string) (*model.Payment, error)
}

type PaymentRepositoryImpl struct {
	db *mongo.Database
}

func NewPaymentRepositoryImpl(db *mongo.Database) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (u *PaymentRepositoryImpl) Create(sid string) (*model.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(STORE_DB_NAME)

	bsid, err := bson.ObjectIDFromHex(sid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	payment := &model.Payment{
		StoreId: bsid,
		Amount: 1000,
	}

	inserted, err := col.InsertOne(ctx, payment)
	if err != nil { 
		return nil, err
	}	

	payment.ID = inserted.InsertedID.(bson.ObjectID)

	fmt.Printf("%v\n", payment)
	return payment, nil
}
