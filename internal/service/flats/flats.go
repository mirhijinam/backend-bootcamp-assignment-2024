package flats

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

type flatRepo interface {
	Create(ctx context.Context, flat dto.NewFlat) (models.Flat, error)
}

type Service struct {
	flatRepo flatRepo
	logger   *zap.Logger
}

func New(hr flatRepo, logger *zap.Logger) Service {
	return Service{
		flatRepo: hr,
		logger:   logger,
	}
}

func (s Service) Create(
	ctx context.Context,
	draftFlat dto.NewFlat,
) (models.Flat, error) {
	return s.flatRepo.Create(ctx, draftFlat)
}
