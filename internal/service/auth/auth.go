package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/pkg/password"
	"go.uber.org/zap"
)

type authRepo interface {
	RegisterUser(ctx context.Context, flat models.User) (uuid.UUID, error)
}

type Service struct {
	authRepo authRepo
	logger   *zap.Logger
}

func New(hr authRepo, logger *zap.Logger) Service {
	return Service{
		authRepo: hr,
		logger:   logger,
	}
}

func (s Service) RegisterUser(
	ctx context.Context,
	draftUser dto.NewUser,
) (uuid.UUID, error) {

	hashed, err := password.HashPassword(draftUser.Password)
	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return uuid.Nil, err
	}

	var user models.User
	user.ID = uuid.New()
	user.Email = draftUser.Email
	user.Password = hashed
	user.Role = draftUser.Role

	return s.authRepo.RegisterUser(ctx, user)
}
