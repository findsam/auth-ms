package util

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/findsam/auth-micro/pkg/config"
	"github.com/findsam/auth-micro/pkg/util"
	"github.com/golang-jwt/jwt/v5"
)

func generateAccessToken() (string, error) {
	claims := jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Envs.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func generateOpaqueToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	str := hex.EncodeToString(bytes)
	return str, nil
}

func GenerateTokens() (*util.TokenPair, error) {
	pair := new(util.TokenPair)

	accessToken, err := generateAccessToken()
	if err != nil {
		return nil, err
	}
	pair.AccessToken = accessToken
	refreshToken, err := generateOpaqueToken()
	if err != nil {
		return nil, err
	}
	pair.RefreshToken = refreshToken
	return pair, nil
}
