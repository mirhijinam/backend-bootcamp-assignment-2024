package flats

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
)

type flatRepo interface {
	Create(ctx context.Context, flat dto.NewFlat) (models.Flat, error)
}

type Service struct {
	flatRepo flatRepo
}

func New(hr flatRepo) Service {
	return Service{
		flatRepo: hr,
	}
}

func (s Service) Create(
	ctx context.Context,
	draftFlat dto.NewFlat,
) (models.Flat, error) {
	return s.flatRepo.Create(ctx, draftFlat)
}
