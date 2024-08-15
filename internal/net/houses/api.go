package houses

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

type service interface {
	Create(ctx context.Context, house dto.NewHouse) (models.House, error)
	GetFlatsByHouseId(ctx context.Context, houseId int, role models.Role) ([]models.Flat, error)
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
