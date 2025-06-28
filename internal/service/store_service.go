package service

import (
	"fmt"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
)

type StoreService struct {
	repo repo.StoreRepository
}

func NewStoreService(repo repo.StoreRepository) *StoreService {
	return &StoreService{repo: repo}
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

func (s *StoreService) GetByUsername(username string) (*model.UserStoreResult, error) {
	store, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return store, nil
}
