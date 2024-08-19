package flat

import (
	"context"
	"errors"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/builder/mapper"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/middleware"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/flats"
)

func (api API) FlatUpdatePost(
	ctx context.Context,
	req generated.OptFlatUpdatePostReq,
) (r generated.FlatUpdatePostRes, _ error) {
	user, ok := ctx.Value(middleware.User).(models.UserClaims)
	if !ok {
		return nil, errors.New("failed to get user from context")
	}

	if user.Role != models.RoleModerator {
		return &generated.R401{}, nil
	}

	flat, err := api.service.UpdateStatus(ctx, dto.FlatUpdateParams{
		Number:      int(req.Value.ID),
		HouseID:     int(req.Value.HouseID),
		Status:      mapper.ReqStatus[req.Value.Status],
		ModeratorID: user.ID,
	})
	if err != nil {
		if errors.Is(err, flats.ErrAlreadyModerated) {
			return &generated.R400{
				Message: "Квартира уже прошла модерацию",
			}, nil
		}

		if errors.Is(err, flats.ErrAnotherModerating) {
			return &generated.R400{
				Message: "Квартира уже на модерации",
			}, nil
		}

		if errors.Is(err, flats.ErrInvalidTransition) {
			return &generated.R400{
				Message: "Невалидный статус перехода",
			}, nil
		}

		return &generated.R5xx{
			Message: err.Error(),
		}, nil
	}

	return &generated.Flat{
		ID:      generated.FlatId(flat.Number),
		HouseID: generated.HouseId(flat.HouseID),
		Rooms:   generated.Rooms(flat.Rooms),
		Price:   generated.Price(flat.Price),
		Status:  mapper.RespStatus[flat.Status],
	}, nil
}
