package flats

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

var (
	ErrAnotherModerating = errors.New("another moderating")
	ErrAlreadyModerated  = errors.New("already moderated")
	ErrInvalidTransition = errors.New("invalid flat status transition")
)

type transactor interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) (err error)
}

type flatsRepo interface {
	Create(ctx context.Context, flat dto.NewFlat) (models.Flat, error)
	Get(ctx context.Context, number int, houseId int) (models.Flat, error)
	UpdateStatus(ctx context.Context, params dto.FlatUpdateParams) (models.Status, uuid.UUID, error)
}

type housesRepo interface {
	Update(ctx context.Context) error
	GetHouseById(ctx context.Context, houseId int) (models.House, error)
}

type Service struct {
	flatsRepo  flatsRepo
	housesRepo housesRepo
	logger     *zap.Logger
	transactor transactor
}

func New(fr flatsRepo, hr housesRepo, logger *zap.Logger, transactor transactor) Service {
	return Service{
		flatsRepo:  fr,
		housesRepo: hr,
		logger:     logger,
		transactor: transactor,
	}
}

func (s Service) Create(
	ctx context.Context,
	draftFlat dto.NewFlat,
) (models.Flat, error) {
	if _, err := s.housesRepo.GetHouseById(ctx, draftFlat.HouseID); err != nil {
		return models.Flat{}, err
	}

	var flat models.Flat
	err := s.transactor.Do(ctx, func(ctx context.Context) error {
		f, err := s.flatsRepo.Create(ctx, draftFlat)
		if err != nil {
			return err
		}

		flat = f

		err = s.housesRepo.Update(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return models.Flat{}, err
	}

	return flat, nil
}

func (s Service) UpdateStatus(
	ctx context.Context,
	params dto.FlatUpdateParams,
) (models.Flat, error) {
	var updatedFlat models.Flat

	err := s.transactor.Do(ctx, func(ctx context.Context) error {
		flat, err := s.flatsRepo.Get(ctx, params.Number, params.HouseID)
		if err != nil {
			return err
		}

		if flat.Status == models.StatusApproved || flat.Status == models.StatusDeclined {
			return ErrAlreadyModerated
		}

		if flat.Status == models.StatusOnModeration && (flat.ModeratorID != nil && *flat.ModeratorID != params.ModeratorID) {
			return ErrAnotherModerating
		}

		if !CanStatusTransition(flat.Status, params.Status) {
			return ErrInvalidTransition
		}

		status, moderatorId, err := s.flatsRepo.UpdateStatus(ctx, params)
		if err != nil {
			return err
		}

		updatedFlat = models.Flat{
			Number:      flat.Number,
			HouseID:     flat.HouseID,
			Rooms:       flat.Rooms,
			Price:       flat.Price,
			Status:      status,
			ModeratorID: &moderatorId,
		}

		return nil
	})

	return updatedFlat, err
}

func CanStatusTransition(prev, next models.Status) bool {
	switch {
	case prev == models.StatusCreated && next == models.StatusOnModeration:
		return true
	case prev == models.StatusOnModeration && next == models.StatusDeclined:
		return true
	case prev == models.StatusOnModeration && next == models.StatusApproved:
		return true
	}

	return false
}
