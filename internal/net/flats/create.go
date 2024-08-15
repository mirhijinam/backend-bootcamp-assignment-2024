package flat

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

// todo: create const
var statusMapping = map[string]generated.Status{
	"created":       generated.StatusCreated,
	"on_moderation": generated.StatusOnModeration,
	"approved":      generated.StatusApproved,
	"declined":      generated.StatusDeclined,
}

func (api API) FlatCreatePost(
	ctx context.Context,
	req generated.OptFlatCreatePostReq,
) (r generated.FlatCreatePostRes, _ error) {
	flat, err := api.service.Create(ctx, dto.NewFlat{
		HouseID: int(req.Value.HouseID),
		Price:   int(req.Value.Price),
		Rooms:   int(req.Value.Rooms),
	})
	if err != nil {
		api.logger.Error("failed to create flat", zap.Error(err))
		return nil, err
	}

	return &generated.Flat{
		ID:      generated.FlatId(flat.Number),
		HouseID: generated.HouseId(flat.HouseID),
		Price:   generated.Price(flat.Price),
		Rooms:   generated.Rooms(flat.Rooms),
		Status:  statusMapping[flat.Status],
	}, nil
}
