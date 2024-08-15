package flat

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

type service interface {
	Create(ctx context.Context, draftFlat dto.NewFlat) (models.Flat, error)
}

type API struct {
	service service
	logger  *zap.Logger
}

func New(service service, logger *zap.Logger) API {
	return API{
		service: service,
		logger:  logger,
	}
}
