package service

import (
	"fmt"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
)

type PaymentService struct {
	repo  repo.PaymentRepository
	store repo.StoreRepository
}

func NewPaymentService(repo repo.PaymentRepository, store repo.StoreRepository) *PaymentService {
	return &PaymentService{repo: repo, store: store}
}

func (s *PaymentService) Create(m *model.CreatePaymentBody) (*model.Payment, error) {
	store, err := s.store.GetById(m.OwnerId)
	if err != nil {
		return nil, err
	}
	tierLen := len(*store.Tiers)

	if tierLen == 0 {
		return nil, fmt.Errorf("store has no tiers")
	}

	tier := (*store.Tiers)[m.Sub]

	payment, err := s.repo.Create(m.OwnerId, tier.Amount)
	if err != nil {
		return nil, err
	}
	return payment, nil
}
