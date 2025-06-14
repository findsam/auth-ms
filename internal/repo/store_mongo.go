package repo

import (
	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const STORE_DB_NAME = "store"

type StoreRepository interface {
	Create() (*model.User, error)
}

type StoreRepositoryImpl struct {
	db *mongo.Database
}

func NewStoreRepositoryImpl(db *mongo.Database) *StoreRepositoryImpl {
	return &StoreRepositoryImpl{
		db: db,
	}
}

func (u *StoreRepositoryImpl) Create() (*model.User, error) {
	return nil, nil
}
