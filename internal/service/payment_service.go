package service

import "github.com/findsam/auth-micro/internal/repo"

type PaymentService struct {
	repo repo.PaymentRepository
}

func NewPaymentService(repo repo.PaymentRepository) *PaymentService {
 return &PaymentService{repo: repo}
}
