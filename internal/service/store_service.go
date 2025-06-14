package service

import (
	"github.com/findsam/auth-micro/internal/repo"
)

type StoreService struct {
	repo repo.StoreRepository
}

func NewStoreService(repo repo.StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) Create(){
}
