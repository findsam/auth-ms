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
	GetById(id string) (*model.PaymentAggregateResult, error)
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
		StripeID: strid,
		Meta:     model.NewMeta(),
	}

	inserted, err := col.InsertOne(ctx, payment)
	if err != nil {
		return nil, err
	}

	payment.ID = inserted.InsertedID.(bson.ObjectID)
	return payment, nil
}


func (u *PaymentRepositoryImpl) GetById(id string) (*model.PaymentAggregateResult, error) {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: objID}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "store"},
			{Key: "localField", Value: "store_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "store"},
		}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$store"}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "store.owner_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user"},
		}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$user"}}}},
		{{Key: "$project", Value: bson.D{
			{Key: "payment", Value: "$$ROOT"},
			{Key: "user", Value: "$user"},
			{Key: "store", Value: "$store"},
		}}},
	}

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    col := u.db.Collection(PAYMENT_DB_NAME)
    cursor, err := col.Aggregate(ctx, pipeline)
    if err != nil {
        return nil, fmt.Errorf("failed to aggregate payments: %w", err)
    }
    defer cursor.Close(ctx)
	var result model.PaymentAggregateResult
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode payment: %w", err)
		}
	}
    fmt.Printf("Aggregate result: %+v\n", result)
    return &result, nil
}