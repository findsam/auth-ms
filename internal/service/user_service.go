package service

import (
	"fmt"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
	util "github.com/findsam/auth-micro/pkg/token"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(u *model.User) (*model.User, error) {



	_, err := s.GetByEmail(u.Email);
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	_, err = util.GenerateTokens()
	if err != nil {
		return nil, fmt.Errorf("failed to generate token pair: %w", err)
	}

	// fmt.Println("Access Token:", tokens.AccessToken)
	// fmt.Println("Refresh Token:", tokens.RefreshToken)

	// user := s.repo.CreateUser()

	return s.repo.CreateUser(u)
}

func (s *UserService) GetByEmail(e string) (*model.User, error) {
	return s.repo.GetByEmail(e)
}
