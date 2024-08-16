package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

type service interface {
	RegisterUser(ctx context.Context, draftUser dto.NewUser) (uuid.UUID, error)
}

type API struct {
	logger  *zap.Logger
	service service
}

func New(service service, logger *zap.Logger) API {
	return API{
		service: service,
		logger:  logger,
	}
}
