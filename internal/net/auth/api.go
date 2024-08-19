package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

type service interface {
	RegisterUser(ctx context.Context, draftUser dto.NewUser) (uuid.UUID, error)
	LoginUser(ctx context.Context, userId uuid.UUID, userPass string) (models.User, error)
}

type API struct {
	logger    *zap.Logger
	service   service
	secretKey string
}

func New(service service, logger *zap.Logger, secretKey string) API {
	return API{
		service:   service,
		logger:    logger,
		secretKey: secretKey,
	}
}
