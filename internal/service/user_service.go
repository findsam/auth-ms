package service

import (
	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/pkg/token"
	"github.com/findsam/auth-micro/pkg/util"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}


func (s *UserService) CreateUser(u *model.User,) (*model.User, *util.TokenPair, error) {
	user, err := s.repo.CreateUser(u)
	if err != nil {
		return nil, nil, err
	}
	tokens, err := token.GenerateTokens(user.ID.Hex())
	if err != nil {
		return nil, nil, err
	}
	return user, tokens, nil
	
}

func (s *UserService) GetByEmail(e string) (*model.User, error) {
	return s.repo.GetByEmail(e)
}
