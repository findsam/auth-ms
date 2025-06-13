package service

import (
	"fmt"

	"github.com/findsam/auth-micro/internal/model"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/pkg/bcrypt"
	"github.com/findsam/auth-micro/pkg/token"
	"github.com/findsam/auth-micro/pkg/util"
)

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(u *model.User) (*model.User, *util.TokenPair, error) {
	pwd, err := bcrypt.HashPassword(u.Password)
	if err != nil {
		return nil, nil, err
	}
	u.Password = pwd
	user, err := s.repo.SignUp(u)
	if err != nil {
		return nil, nil, err
	}
	tokens, err := token.GenerateTokens(user.ID.Hex())
	if err != nil {
		return nil, nil, err
	}
	return user, tokens, nil
}

func (s *UserService) SignIn(u *model.UserSignInRequest) (*model.User, *util.TokenPair, error) {

	user, err := s.repo.GetByEmail(u.Email)
	if err != nil {
		return nil, nil, err
	}

	if !bcrypt.ComparePasswords(user.Password, u.Password) {
		return nil, nil, fmt.Errorf("invalid credentials")
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

func (s *UserService) GetById(id string) (*model.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) GetByUsername(id string) (*model.User, error) {
	return s.repo.GetByUsername(id)
}

func (s *UserService) Me(t string) (*model.User, error) {
	return nil, nil
}

func (s *UserService) Refresh(uid string) (*util.TokenPair, error) {
	tokens, err := token.GenerateTokens(uid)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}
	return tokens, nil
}
