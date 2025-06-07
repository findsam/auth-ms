package repo

import (
	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository interface {
	GetByID(id string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *mongo.Client
}

func NewUserRepositoryImpl(db *mongo.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) GetByID(id string) (*model.User, error) {
	return nil, nil
}