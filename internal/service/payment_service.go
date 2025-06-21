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


func (s *PaymentService) Create(*model.Payment) (*model.Payment, error) {
	payment, err := s.repo.Create("68568115de41192bde902111")
	if err != nil {
		return nil, err
	}
	return payment, nil
}
