package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/findsam/auth-micro/internal/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *mongo.Database
}

func NewUserRepositoryImpl(db *mongo.Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection("users")

	insert, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	oid, ok := insert.InsertedID.(bson.ObjectID)
	if !ok {
		return nil, fmt.Errorf("error while creating User %v", err)
	}
	user.ID = oid.Hex()
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection("users")

	user := new(model.User)
	err := col.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
