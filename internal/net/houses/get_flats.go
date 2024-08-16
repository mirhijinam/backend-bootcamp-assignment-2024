package houses

import (
	"context"
	"errors"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/middleware"
	"go.uber.org/zap"
)

func (api API) HouseIDGet(
	ctx context.Context,
	params generated.HouseIDGetParams,
) (r generated.HouseIDGetRes, _ error) {
	user, ok := ctx.Value(middleware.User).(models.UserClaims)
	if !ok {
		return nil, errors.New("failed to get user from context")
	}

	api.logger.Info("flat list requester (info from context):", zap.Any("userID", user.ID), zap.Any("userRole", user.Role))

	intermediateFlats, err := api.service.GetFlatsByHouseId(ctx, int(params.ID), models.Role(user.Role))
	if err != nil {
		api.logger.Error("failed to get flats by the houseId", zap.Error(err))
		return nil, err
	}

	flats := make([]generated.Flat, 0, len(intermediateFlats))
	for _, flat := range intermediateFlats {
		flats = append(flats, generated.Flat{
			ID:      generated.FlatId(flat.Number),
			HouseID: generated.HouseId(flat.HouseID),
			Price:   generated.Price(flat.Price),
			Rooms:   generated.Rooms(flat.Rooms),
			Status:  generated.Status(flat.Status),
		})
	}

	return &generated.HouseIDGetOK{Flats: flats}, nil
}
