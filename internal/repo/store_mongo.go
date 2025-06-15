package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const STORE_DB_NAME = "store"

type StoreRepository interface {
	Create(oid string) (*model.Store, error)
	GetById(oid string) (*model.Store, error)
}

type StoreRepositoryImpl struct {
	db *mongo.Database
}

func NewStoreRepositoryImpl(db *mongo.Database) *StoreRepositoryImpl {
	return &StoreRepositoryImpl{
		db: db,
	}
}

func (u *StoreRepositoryImpl) Create(oid string) (*model.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(STORE_DB_NAME)

	ownerID, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	store := &model.Store{
		OwnerId:     bson.ObjectID(ownerID),
		Name:        "Default Store",
		Description: "This is a default store",
	}

	_, err = col.InsertOne(ctx, store)

	if err != nil {
		return nil, err
	}

	return store, nil
}


func (u *StoreRepositoryImpl) GetById(oid string) (*model.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(STORE_DB_NAME)

	ownerID, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}
	store := &model.Store{}

	err = col.FindOne(
		ctx,
		bson.M{"owner_id": bson.ObjectID(ownerID)},
	).Decode(store)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("store not found for owner ID: %s", oid)
		}
		return nil, err
	}

	fmt.Printf("Store: %+v\n", store)

	return store, nil
}