package houses

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
)

type houseRepo interface {
	Create(ctx context.Context, house dto.NewHouse) (models.House, error)
	GetFlatsByHouseId(ctx context.Context, houseId int, role models.Role) ([]models.Flat, error)
}

type Service struct {
	houseRepo houseRepo
}

func New(hr houseRepo) Service {
	return Service{
		houseRepo: hr,
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
