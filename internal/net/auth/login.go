package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/auth"
	"go.uber.org/zap"
)

func (api API) LoginPost(
	ctx context.Context,
	req generated.OptLoginPostReq,
) (r generated.LoginPostRes, _ error) {
	user, err := api.service.LoginUser(ctx, uuid.UUID(req.Value.ID.Value), string(req.Value.Password.Value))
	if err != nil {
		if errors.Is(err, auth.ErrWrongPassword) {
			return &generated.LoginPostNotFound{}, nil
		}
		api.logger.Error("failed to register user", zap.Error(err))
		return &generated.LoginPostBadRequest{}, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		},
		User: models.UserClaims{
			ID:   user.ID,
			Role: user.Role,
		},
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		api.logger.Error("failed to create token string", zap.Error(err))
		return nil, err
	}

	return &generated.LoginPostOK{
		Token: generated.OptToken{
			Value: generated.Token(tokenString),
			Set:   true,
		},
	}, nil
}
