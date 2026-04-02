package jwtmanager

import (
	"errors"
	"fitfeed/auth/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func New(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey, tokenDuration: tokenDuration}
}

type UserClaims struct {
	jwt.RegisteredClaims
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (m *JWTManager) GenerateToken(user entity.User) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.tokenDuration)),
		},
		ID:       user.ID.String(),
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) ValidateToken(tokenStr string) (*entity.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}
			return []byte(m.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// We can convert claims back to entity.UserClaims here or change the interface
	return &entity.UserClaims{
		Username: claims.Username,
	}, nil
}
