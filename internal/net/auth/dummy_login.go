package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"go.uber.org/zap"
)

const expirationTime = 12 * time.Hour

func (api API) DummyLoginGet(
	ctx context.Context,
	params generated.DummyLoginGetParams,
) (r generated.DummyLoginGetRes, _ error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		},
		User: models.UserClaims{
			ID:   uuid.NewString(),
			Role: models.Role(params.UserType),
		},
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		api.logger.Error("failed to create token string", zap.Error(err))
		return nil, err
	}

	return &generated.DummyLoginGetOK{
		Token: generated.OptToken{
			Value: generated.Token(tokenString),
			Set:   true,
		},
	}, nil
}
