package middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"go.uber.org/zap"
)

type userKey string

const User userKey = "user"

type Middleware struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Middleware {
	return Middleware{
		logger: logger,
	}
}

func (m Middleware) HandleBearerAuth(
	ctx context.Context,
	operationName string,
	t generated.BearerAuth,
) (context.Context, error) {
	m.logger.Info("TOKEN:", zap.Any("token", t.Token))

	token, err := jwt.ParseWithClaims(t.GetToken(), &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil // todo: add secret in config
	})

	if err != nil {
		m.logger.Error("failed to parse token", zap.Error(err))
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Если токен валиден, добавляем информацию о пользователе в контекст
	ctx = context.WithValue(ctx,
		User,
		models.UserClaims{
			ID:   claims.User.ID,
			Role: claims.User.Role,
		})

	return ctx, err
}
