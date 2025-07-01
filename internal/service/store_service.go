package service

import (
	"fmt"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
)

type StoreService struct {
	repo repo.StoreRepository
	user repo.UserRepository
}

func NewStoreService(repo repo.StoreRepository, user repo.UserRepository) *StoreService {
	return &StoreService{repo: repo, user: user}
}

func (s *StoreService) Create(oid string) (*model.Store, error) {
	store, err := s.repo.GetById(oid)
	if err != nil && err.Error() != "store not found" {
		return nil, err
	}

	if store != nil {
		return nil, fmt.Errorf("store already exists")
	}

	store, err = s.repo.Create(oid)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *StoreService) GetByUsername(username string) (map[string]interface{}, error) {
	user, err := s.user.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	store, err := s.repo.GetById(user.Id.Hex())
	if err != nil {
		return nil, fmt.Errorf("failed to get store by user id: %w", err)
	}
	return map[string]interface{}{
		"store": store,
		"user":  user,
	}, nil
}
