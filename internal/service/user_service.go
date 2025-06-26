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

func (s *UserService) SignUp(u *model.User) (*model.UserPublic, *util.TokenPair, error) {
	user, err := s.repo.GetByEmail(u.Email)
	if err != nil && err.Error() != "user not found" {
		return nil, nil, err
	}

	if user != nil {
		return nil, nil, fmt.Errorf("user already exists")
	}

	pwd, err := bcrypt.HashPassword(u.Password)
	if err != nil {
		return nil, nil, err
	}

	u.Password = pwd
	u.ToDatabase()

	user, err = s.repo.SignUp(u)
	if err != nil {
		return nil, nil, err
	}

	tokens, err := token.GenerateTokens(user.Id.Hex())
	if err != nil {
		return nil, nil, err
	}
	return user.ToPublic(), tokens, nil
}

func (s *UserService) SignIn(u *model.UserSignInRequest) (*model.UserPublic, *util.TokenPair, error) {
	user, err := s.repo.GetByEmail(u.Email)
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, fmt.Errorf("no user found")
	}

	if !bcrypt.ComparePasswords(user.Password, u.Password) {
		return nil, nil, fmt.Errorf("invalid credentials")
	}

	tokens, err := token.GenerateTokens(user.Id.Hex())
	if err != nil {
		return nil, nil, err
	}
	return user.ToPublic(), tokens, nil
}

func (s *UserService) GetByEmail(e string) (*model.UserPublic, error) {
	user, err := s.repo.GetByEmail(e)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user.ToPublic(), nil
}

func (s *UserService) GetById(id string) (*model.UserPublic, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user.ToPublic(), nil
}

func (s *UserService) GetByUsername(username string) (*model.UserPublic, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	return user.ToPublic(), nil
}

func (s *UserService) Me(t string) (*model.UserPublic, error) {
	return nil, nil
}

func (s *UserService) Refresh(uid string) (*util.TokenPair, error) {
	tokens, err := token.GenerateTokens(uid)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}
	return tokens, nil
}
