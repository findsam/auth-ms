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
	Create(sid string, amount float64) (*model.Payment, error)
	GetByStoreId(sid string) ([]*model.Payment, error)
}

type PaymentRepositoryImpl struct {
	db *mongo.Database
}

func NewPaymentRepositoryImpl(db *mongo.Database) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (u *PaymentRepositoryImpl) Create(sid string, amount float64) (*model.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(PAYMENT_DB_NAME)

	bsid, err := bson.ObjectIDFromHex(sid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	payment := &model.Payment{
		StoreId: bsid,
		Amount:  amount,
		Meta:    model.NewMeta(),
	}

	inserted, err := col.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}

	payment.ID = inserted.InsertedID.(bson.ObjectID)
	return payment, nil
}

func (u *PaymentRepositoryImpl) GetByStoreId(sid string) ([]*model.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bsid, err := bson.ObjectIDFromHex(sid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	col := u.db.Collection(PAYMENT_DB_NAME)
	filter := bson.M{"store_id": bsid}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find payments: %w", err)
	}
	defer cursor.Close(ctx)

	var payments []*model.Payment
	if err = cursor.All(ctx, &payments); err != nil {
		return nil, fmt.Errorf("failed to decode documents: %w", err)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor iteration error: %w", err)
	}

	fmt.Printf("Found %d payments for store ID: %s\n", len(payments), sid)

	return payments, nil
}
