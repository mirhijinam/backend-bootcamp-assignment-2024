package auth

import (
	"context"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

func (api API) RegisterPost(
	ctx context.Context,
	req generated.OptRegisterPostReq,
) (r generated.RegisterPostRes, _ error) {
	userID, err := api.service.RegisterUser(ctx, dto.NewUser{
		Email:    string(req.Value.Email),
		Password: string(req.Value.Password),
		Role:     models.Role(req.Value.UserType),
	})
	if err != nil {
		api.logger.Error("failed to register user", zap.Error(err))
		return &generated.RegisterPostBadRequest{}, err
	}

	api.logger.Info("registered user", zap.Any("id", userID))

	return &generated.RegisterPostOK{
		UserID: generated.OptUserId{
			Value: generated.UserId(userID),
			Set:   true,
		},
	}, nil
}
