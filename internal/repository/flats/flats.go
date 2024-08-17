package flats

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
	"go.uber.org/zap"
)

var ErrFlatNumberAlreadyExists = errors.New("flat number already exists")

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

func (r Repo) Create(
	ctx context.Context,
	draftFlat dto.NewFlat,
) (models.Flat, error) {
	const op = `repo.Flats.Create`

	// todo: удостовериться, что у нас не будет состояния гонки при попытке задать номер для квартиры

	builder := r.builder.
		Insert("flats").
		Columns(
			"flat_number",
			"house_id",
			"price",
			"rooms",
		).
		Values(
			draftFlat.Number,
			draftFlat.HouseID,
			draftFlat.Price,
			draftFlat.Rooms,
		).
		Suffix(
			"RETURNING flat_number, house_id, price, rooms, status",
		)

	sql, args, err := builder.ToSql()
	if err != nil {
		return models.Flat{}, fmt.Errorf("%s: %w", op, err)
	}

	r.logger.Info(op, zap.Any("sql request", sql))

	var flat models.Flat
	err = r.pool.QueryRow(ctx, sql, args...).Scan(
		&flat.Number,
		&flat.HouseID,
		&flat.Price,
		&flat.Rooms,
		&flat.Status,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return models.Flat{}, ErrFlatNumberAlreadyExists
			}
		}

		return models.Flat{}, err
	}

	return flat, nil
}

func (r Repo) UpdateStatus(
	ctx context.Context,
	params dto.FlatUpdateParams,
) (models.Status, uuid.UUID, error) {
	const op = `repo.Flats.UpdateStatus`

	builder := r.builder.
		Update("flats").
		Set("status", string(params.Status)).
		Set("moderator_id", params.ModeratorID).
		Where(sq.Eq{
			"flat_number": params.Number,
			"house_id":    params.HouseID,
		}).
		Suffix(
			"RETURNING status, moderator_id",
		)

	sql, args, err := builder.ToSql()
	if err != nil {
		return "", uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	var (
		status      models.Status
		moderatorId uuid.UUID
	)

	err = r.pool.QueryRow(ctx, sql, args...).Scan(&status, &moderatorId)
	if err != nil {
		return "", uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return status, moderatorId, nil
}

func (r Repo) Get(
	ctx context.Context,
	number int,
	houseId int,
) (models.Flat, error) {
	const op = `repo.Flats.Get`

	builder := r.builder.
		Select(
			"flat_number",
			"house_id",
			"price",
			"rooms",
			"status",
			"moderator_id",
		).
		From("flats").
		Where(sq.Eq{
			"house_id":    houseId,
			"flat_number": number,
		})

	sql, args, err := builder.ToSql()
	if err != nil {
		return models.Flat{}, fmt.Errorf("%s: %w", op, err)
	}

	var flat models.Flat
	err = r.pool.QueryRow(ctx, sql, args...).Scan(
		&flat.Number,
		&flat.HouseID,
		&flat.Price,
		&flat.Rooms,
		&flat.Status,
		&flat.ModeratorID,
	)
	if err != nil {
		return models.Flat{}, fmt.Errorf("%s: %w", op, err)
	}

	return flat, nil
}
