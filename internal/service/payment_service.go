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
}

func NewPaymentService(repo repo.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) GetById(username string, id string) (any, error) {
	result, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment by id: %w", err)
	}
	if result.User.Username != username {
		return nil, fmt.Errorf("user %s does not own payment %s", username, id)
	}
	stripe.Key = config.Envs.STRIPE_PWD
	resi, err := paymentintent.Get(result.Payment.StripeId, &stripe.PaymentIntentParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to get payment intent: %w", err)
	}
	fmt.Printf("%v\n", resi)
	return map[string]interface{}{
		"payment": result,
		"intent":  resi,
	}, nil
}

// func (s *PaymentService) Create(m *model.CreatePaymentBody) (*model.Payment, error) {
// 	stripe.Key = config.Envs.STRIPE_PWD
// 	store, err := s.store.GetByStoreId(m.StoreId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	tierLen := len(*store.Tiers)

// 	if tierLen == 0 {
// 		return nil, fmt.Errorf("store has no tiers")
// 	}

// 	tier := (*store.Tiers)[m.Tier]

// 	params := &stripe.PaymentIntentParams{
// 		Amount:   stripe.Int64(int64(float64(tier.Amount) * 1.10)),
// 		Currency: stripe.String(stripe.CurrencyUSD),
// 		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
// 			Enabled: stripe.Bool(true),
// 		},
// 	}

// 	result, err := paymentintent.New(params)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create payment intent: %w", err)
// 	}

// 	payment, err := s.repo.Create(m.StoreId, result.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return payment, nil
// }

// func (s *PaymentService) GetById(username string, id string) (*model.Payment, *stripe.PaymentIntent, error) {
// 	user, err := s.user.GetByUsername(username)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to get user by id: %w", err)
// 	}
// 	payment, err := s.repo.GetById(id)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to get payment by id: %w", err)
// 	}

// 	store, err := s.store.GetByStoreId(payment.StoreId.Hex())
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("failed to get store by id: %w", err)
// 	}

// 	if user.ID != store.OwnerId {
// 		return nil, nil, fmt.Errorf("user %s is not the owner of store %s", username, store.ID)
// 	}

// 	stripe.Key = config.Envs.STRIPE_PWD
// 	result, err := paymentintent.Get(payment.StripeID, &stripe.PaymentIntentParams{})
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("intent not found: %w", err)
// 	}

// 	return payment, result, nil
// }
