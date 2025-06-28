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
	GetById(id string) (*model.PaymentResponse, error)
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

func (u *PaymentRepositoryImpl) GetById(id string) (*model.PaymentResponse, error) {
	pid, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: pid}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "store"},
			{Key: "localField", Value: "store_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "store"},
		}}},
		{{Key: "$unwind", Value: "$store"}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "store.owner_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "user"},
		}}},
		{{Key: "$unwind", Value: "$user"}},
		{{Key: "$project", Value: bson.D{
			{Key: "payment", Value: "$$ROOT"},
			{Key: "store", Value: "$store"},
			{Key: "user", Value: "$user"},
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

	var result model.PaymentResponse
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode result: %w", err)
		}
	}

	return &result, nil
}

// func (u *StoreRepositoryImpl) GetByUsername(username string) (*model.UserStoreResult, error) {
// pipeline := mongo.Pipeline{
// 	{{Key: "$lookup", Value: bson.D{
// 		{Key: "from", Value: "users"},
// 		{Key: "pipeline", Value: bson.A{
// 			bson.D{{Key: "$match", Value: bson.D{{Key: "username", Value: username}}}},
// 		}},
// 		{Key: "as", Value: "user"},
// 	}}},
// 	{{Key: "$unwind", Value: "$user"}},
// 	{{Key: "$lookup", Value: bson.D{
// 		{Key: "from", Value: "store"},
// 		{Key: "localField", Value: "user._id"},
// 		{Key: "foreignField", Value: "owner_id"},
// 		{Key: "as", Value: "store"},
// 	}}},
// 	{{Key: "$unwind", Value: "$store"}},
// 	{{Key: "$project", Value: bson.D{
// 		{Key: "store", Value: "$store"},
// 		{Key: "user", Value: "$user"},
// 	}}},
// }

// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// col := u.db.Collection(PAYMENT_DB_NAME)
// cursor, err := col.Aggregate(ctx, pipeline)
// if err != nil {
// 	return nil, fmt.Errorf("failed to aggregate payments: %w", err)
// }
// defer cursor.Close(ctx)

// var result model.UserStoreResult
// if cursor.Next(ctx) {
// 	if err := cursor.Decode(&result); err != nil {
// 		return nil, fmt.Errorf("failed to decode result: %w", err)
// 	}
// }

// 	return &result, nil
// }
