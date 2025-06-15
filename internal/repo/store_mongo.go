package repo

import (
	"context"
	"time"

	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const STORE_DB_NAME = "store"

type StoreRepository interface {
	Create() (*model.Store, error)
}

type StoreRepositoryImpl struct {
	db *mongo.Database
}

func NewStoreRepositoryImpl(db *mongo.Database) *StoreRepositoryImpl {
	return &StoreRepositoryImpl{
		db: db,
	}
}

func (u *StoreRepositoryImpl) Create() (*model.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection(STORE_DB_NAME)

	store := &model.Store{
		Name:        "Default Store",
		Description: "This is a default store",
	}

	_, err := col.InsertOne(ctx, store);

	if err != nil {
		return nil, err
	}

	return store, nil
}
