package service

import (
	"fmt"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/pkg/config"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

type PaymentService struct {
	repo  repo.PaymentRepository
	store repo.StoreRepository
}

func NewPaymentService(repo repo.PaymentRepository, store repo.StoreRepository) *PaymentService {
	return &PaymentService{repo: repo, store: store}
}

func (s *PaymentService) Create(m *model.CreatePaymentBody) (*model.Payment, error) {
	stripe.Key = config.Envs.STRIPE_PWD
	store, err := s.store.GetByStoreId(m.StoreId)
	if err != nil {
		return nil, err
	}
	tierLen := len(*store.Tiers)

	if tierLen == 0 {
		return nil, fmt.Errorf("store has no tiers")
	}

	tier := (*store.Tiers)[m.Sub]

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(float64(tier.Amount) * 1.10)),
		Currency: stripe.String(stripe.CurrencyUSD),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
		  Enabled: stripe.Bool(true),
		},
	  };

	result, err := paymentintent.New(params);

	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}

	payment, err := s.repo.Create(m.StoreId, result.ID, tier.Amount)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) GetById(id string) (*model.Payment, error) {
	payment, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment by id: %w", err)
	}
	stripe.Key = config.Envs.STRIPE_PWD
	
	_, err = paymentintent.Get(payment.StripeID, &stripe.PaymentIntentParams{});
	if err != nil {
		return nil, fmt.Errorf("intent not found: %w", err)
	}	

	return payment, nil
}
