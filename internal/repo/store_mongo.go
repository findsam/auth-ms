package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const STORE_DB_NAME = "store"

type StoreRepository interface {
	Create(oid string) (*model.Store, error)
	GetById(oid string) (*model.Store, error)
	GetByStoreId(oid string) (*model.Store, error)
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

	boid, err := bson.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}
	store := &model.Store{
		OwnerId: boid,
		Tiers: &[]model.Tier{
			{
				Amount:      1000,
				Description: "Basic Tier",
				Benefits:    []string{"Access to basic features", "Email support"},
			},
			{
				Amount:      5000,
				Description: "Second Tier",
				Benefits:    []string{"Access to basic features", "Email support"},
			},
			{
				Amount:      8123,
				Description: "Third Tier",
				Benefits:    []string{"Access to basic features", "Email support"},
			},
			{
				Amount:      100000,
				Description: "Fourth Tier",
				Benefits:    []string{"Access to basic features", "Email support"},
			},
		},
		Meta: model.NewMeta(),
	}

	inserted, err := col.InsertOne(ctx, store)

	if err != nil {
		return nil, err
	}

	store.Id = inserted.InsertedID.(bson.ObjectID)
	return store, nil
}

func (u *StoreRepositoryImpl) GetById(oid string) (*model.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(STORE_DB_NAME)

	boid, err := bson.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	store := &model.Store{}
	err = col.FindOne(
		ctx,
		bson.M{"owner_id": boid},
	).Decode(store)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("store not found")
		}
		return nil, err
	}

	return store, nil
}

func (u *StoreRepositoryImpl) GetByStoreId(oid string) (*model.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(STORE_DB_NAME)

	boid, err := bson.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}

	store := &model.Store{}
	err = col.FindOne(
		ctx,
		bson.M{"_id": boid},
	).Decode(store)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("store not found")
		}
		return nil, err
	}

	return store, nil
}
