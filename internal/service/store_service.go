package service

import (
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
	store, err := s.repo.Create(oid)
	if err != nil {
		return nil, err
	}
	return store, nil
}


func (s *StoreService) GetById(id string) (*model.Store, error) {
	store, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return store, nil
}