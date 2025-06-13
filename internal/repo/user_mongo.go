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

const (
	COLLECTION_NAME = "users"
)

type UserRepository interface {
	SignUp(user *model.User) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetById(id string) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *mongo.Database
}

func NewUserRepositoryImpl(db *mongo.Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) SignUp(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := u.db.Collection("users")

	exists, err := u.GetByEmail(user.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	if exists != nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	_, err = col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	inserted, err := u.GetByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	return inserted, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	col := u.db.Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := new(model.User)
	err := col.FindOne(
		ctx,
		bson.M{"email": email},
	).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetById(id string) (*model.User, error) {
	col := u.db.Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID format: %v", err)
	}
	user := new(model.User)

	err = col.FindOne(
		ctx,
		bson.M{"_id": bson.ObjectID(oid)},
	).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetByUsername(username string) (*model.User, error) {
	col := u.db.Collection(COLLECTION_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user := new(model.User)

	err := col.FindOne(
		ctx,
		bson.M{"username": username},
	).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
