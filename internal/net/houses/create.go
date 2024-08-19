package houses

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/pkg/pointer"
	"go.uber.org/zap"
)

func (api API) HouseCreatePost(
	ctx context.Context,
	req generated.OptHouseCreatePostReq,
) (r generated.HouseCreatePostRes, _ error) {
	var developer *string
	if !req.Value.Developer.IsNull() {
		developer = (*string)(&req.Value.Developer.Value)
	}
	house, err := api.service.Create(ctx, dto.NewHouse{
		Address:   string(req.Value.Address),
		Year:      int(req.Value.Year),
		Developer: developer,
	})

	if err != nil {
		api.logger.Error("failed to create house", zap.Error(err))
		return &generated.R5xx{
			Message: err.Error(),
		}, nil
	}

	return &generated.House{
		ID:      generated.HouseId(house.ID),
		Address: generated.Address(house.Address),
		Year:    generated.Year(house.Year),
		Developer: generated.OptNilDeveloper{
			Value: generated.Developer(pointer.FromPtr(house.Developer)),
			Set:   house.Developer != nil,
			Null:  house.Developer == nil,
		},
		CreatedAt: generated.OptDate{
			Value: generated.Date(house.CreatedAt),
			Set:   true,
		},
		UpdateAt: generated.OptDate{
			Value: generated.Date(pointer.FromPtr(house.UpdateAt)),
			Set:   house.UpdateAt != nil,
		},
	}, nil
}
