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

type PaymentRepository interface {
	Create(sid string, strid string) (*model.Payment, error)
	GetById(id string) (*model.Payment, error)
}

type PaymentRepositoryImpl struct {
	db *mongo.Database
}

func NewPaymentRepositoryImpl(db *mongo.Database) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (u *PaymentRepositoryImpl) Create(sid string, strid string) (*model.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(PAYMENT_DB_NAME)

	bsid, err := bson.ObjectIDFromHex(sid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	payment := &model.Payment{
		StoreId:  bsid,
		StripeId: strid,
		Meta:     model.NewMeta(),
	}

	inserted, err := col.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}

	payment.ID = inserted.InsertedID.(bson.ObjectID)
	return payment, nil
}

func (u *PaymentRepositoryImpl) GetById(id string) (*model.Payment, error) {
	pid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(PAYMENT_DB_NAME)

	var payment model.Payment
	err = col.FindOne(ctx, bson.M{"_id": pid}).Decode(&payment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find payment: %w", err)
	}
	return &payment, nil
}