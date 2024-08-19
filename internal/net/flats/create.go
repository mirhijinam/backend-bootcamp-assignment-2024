package flat

import (
	"context"
	"errors"
	"fmt"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/builder/mapper"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/flats"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/houses"
	"go.uber.org/zap"
)

func (api API) FlatCreatePost(
	ctx context.Context,
	req generated.OptFlatCreatePostReq,
) (r generated.FlatCreatePostRes, _ error) {
	flat, err := api.service.Create(ctx, dto.NewFlat{
		Number:  int(req.Value.Number),
		HouseID: int(req.Value.HouseID),
		Price:   int(req.Value.Price),
		Rooms:   int(req.Value.Rooms),
	})
	if err != nil {
		if errors.Is(err, houses.ErrHouseNotExists) {
			return &generated.R400{
				Message: fmt.Sprintf("Дом с айди: %d не найден", req.Value.HouseID),
			}, nil
		}

		if errors.Is(err, flats.ErrFlatNumberAlreadyExists) {
			return &generated.R400{
				Message: fmt.Sprintf("Номер квартиры: %d уже существует", req.Value.Number),
			}, nil
		}

		api.logger.Error("failed to create flat", zap.Error(err))
		return &generated.R5xx{
			Message: err.Error(),
		}, nil
	}

	return &generated.Flat{
		ID:      generated.FlatId(flat.Number),
		HouseID: generated.HouseId(flat.HouseID),
		Price:   generated.Price(flat.Price),
		Rooms:   generated.Rooms(flat.Rooms),
		Status:  mapper.RespStatus[flat.Status],
	}, nil
}
