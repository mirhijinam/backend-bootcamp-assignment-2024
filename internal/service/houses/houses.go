package houses

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

type houseRepo interface {
	Create(ctx context.Context, house dto.NewHouse) (models.House, error)
	GetFlatsByHouseId(ctx context.Context, houseId int, role models.Role) ([]models.Flat, error)
}

type Service struct {
	houseRepo houseRepo
	logger    *zap.Logger
}

func New(hr houseRepo, logger *zap.Logger) Service {
	return Service{
		houseRepo: hr,
		logger:    logger,
	}
}

func (s Service) Create(
	ctx context.Context,
	house dto.NewHouse,
) (models.House, error) {
	return s.houseRepo.Create(ctx, house)
}

func (s Service) GetFlatsByHouseId(
	ctx context.Context,
	houseId int,
	role models.Role,
) ([]models.Flat, error) {
	return s.houseRepo.GetFlatsByHouseId(ctx, houseId, role)
}
