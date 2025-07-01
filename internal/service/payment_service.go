package service

import (
	"fmt"

	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/pkg/config"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

type PaymentService struct {
	repo repo.PaymentRepository
	store repo.StoreRepository
	user repo.UserRepository
}

func NewPaymentService(repo repo.PaymentRepository, store repo.StoreRepository, user repo.UserRepository) *PaymentService {
	return &PaymentService{repo: repo, store: store, user: user}
}

func (s *PaymentService) GetById(username string, id string) (any, error) {
	payment, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	store, err := s.store.GetByStoreId(payment.StoreId.Hex())

	if err != nil {
		return nil, err
	}

	user, err := s.user.GetById(store.OwnerId.Hex())
	
	if err != nil {
		return nil, err
	}

	if user.Username != username {
		return nil, fmt.Errorf("user %s does not own payment %s", username, id)
	}

	stripe.Key = config.Envs.STRIPE_PWD
	result, err := paymentintent.Get(payment.StripeId, &stripe.PaymentIntentParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to get payment intent: %w", err)
	}

	return map[string]interface{}{
		"payment": payment,
		"store":   store,
		"user":    user,
		"intent":  result,
	}, nil
}

func (s *PaymentService) Create(username string) (any, error) {
	result, err := s.store.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("failed to get store by username: %w", err)
	}

	if len(*result.Store.Tiers) == 0 {
		return nil, fmt.Errorf("store has no tiers")
	}

	payment, err := s.repo.Create()
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}


	return nil, nil
}