package service

import (
	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
)

type PaymentService struct {
	repo repo.PaymentRepository
}

func NewPaymentService(repo repo.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(sid string) (*model.Payment, error) {
	payment, err := s.repo.Create(sid)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) GetByStoreId(sid string) ([]*model.Payment, error) {
	payments, err := s.repo.GetByStoreId(sid)
	if err != nil {
		return nil, err
	}
	return payments, nil
}