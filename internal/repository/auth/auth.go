package auth

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"go.uber.org/zap"
)

type Repo struct {
	pool    *pgxpool.Pool
	builder sq.StatementBuilderType
	logger  *zap.Logger
}

func New(pool *pgxpool.Pool, logger *zap.Logger) Repo {
	return Repo{
		pool:    pool,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		logger:  logger,
	}
}

func (r Repo) RegisterUser(
	ctx context.Context,
	userToAdd models.User,
) (uuid.UUID, error) {
	const op = `repo.Auth.RegisterUser`

	builder := r.builder.
		Insert("users").
		Columns(
			"id",
			"email",
			"password",
			"role",
		).
		Values(
			userToAdd.ID,
			userToAdd.Email,
			userToAdd.Password,
			userToAdd.Role,
		).
		Suffix(
			"RETURNING id",
		)

	sql, args, err := builder.ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	var userID uuid.UUID
	err = r.pool.QueryRow(ctx, sql, args...).Scan(
		&userID,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userID, nil
}
