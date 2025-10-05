package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ValidateToken(tokenStr string) (uuid.UUID, error)
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewService(secret string, duration time.Duration) JWTService {
	return &jwtService{
		secretKey:     secret,
		tokenDuration: duration,
	}
}

func (j *jwtService) GenerateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(j.tokenDuration).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return uuid.Nil, errors.New("token expired")
			}
		}

		idStr, ok := claims["user_id"].(string)
		if !ok {
			return uuid.Nil, errors.New("invalid token payload")
		}
		return uuid.Parse(idStr)
	}

	return uuid.Nil, errors.New("invalid token")
}
