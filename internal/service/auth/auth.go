package auth

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	pass "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/pkg/password"
	"go.uber.org/zap"
)

var ErrWrongPassword = errors.New("wrong password")

type authRepo interface {
	RegisterUser(ctx context.Context, flat models.User) (uuid.UUID, error)
	LoginUser(ctx context.Context, userId uuid.UUID, userPass string) (models.User, error)
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
	hashed, err := pass.HashPassword(draftUser.Password)
	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return uuid.Nil, err
	}

	return s.authRepo.RegisterUser(ctx, models.User{
		ID:       uuid.New(),
		Email:    draftUser.Email,
		Password: hashed,
		Role:     draftUser.Role,
	})
}

func (s Service) LoginUser(
	ctx context.Context,
	userId uuid.UUID,
	userPass string,
) (models.User, error) {
	user, err := s.authRepo.LoginUser(ctx, userId, userPass)
	if err != nil {
		return models.User{}, err
	}

	ok := pass.CheckPassword(user.Password, userPass)
	if !ok {
		return models.User{}, ErrWrongPassword
	}

	return user, nil
}
