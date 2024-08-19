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

var RoleUUID = map[models.Role]uuid.UUID{
	models.RoleModerator: uuid.MustParse("e39286e7-b194-4123-94fb-c69260345f40"),
	models.RoleClient:    uuid.MustParse("ec4eca76-bd8d-4602-98dc-784c568979a8"),
}

const dummyExpirationTime = 365 * (24 * time.Hour)

func (api API) DummyLoginGet(
	ctx context.Context,
	params generated.DummyLoginGetParams,
) (r generated.DummyLoginGetRes, _ error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(dummyExpirationTime)),
		},
		User: models.UserClaims{
			ID:   RoleUUID[models.Role(params.UserType)],
			Role: models.Role(params.UserType),
		},
	})

	tokenString, err := token.SignedString([]byte(api.secretKey))
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
