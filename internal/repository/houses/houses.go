package houses

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
)

var ErrHouseNotExists = errors.New("house not exists")

type PgxIface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type Repo struct {
	pool    PgxIface
	builder sq.StatementBuilderType
}

func New(pool PgxIface) Repo {
	return Repo{
		pool:    pool,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r Repo) Create(
	ctx context.Context,
	draftHouse dto.NewHouse,
) (models.House, error) {
	const op = `repo.House.Create`

	builder := r.builder.
		Insert("houses").
		Columns(
			"address",
			"year",
			"developer",
		).
		Values(
			draftHouse.Address,
			draftHouse.Year,
			draftHouse.Developer,
		).
		Suffix(
			"RETURNING id, address, year, developer, created_at, update_at",
		)

	sql, args, err := builder.ToSql()
	if err != nil {
		return models.House{}, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("%s: sql request = %s", op, sql)

	var house models.House
	err = r.pool.QueryRow(ctx, sql, args...).Scan(
		&house.ID,
		&house.Address,
		&house.Year,
		&house.Developer,
		&house.CreatedAt,
		&house.UpdateAt,
	)
	if err != nil {
		return models.House{}, fmt.Errorf("%s: %w", op, err)
	}

	return house, nil
}

func (r Repo) GetFlatsByHouseId(
	ctx context.Context,
	houseId int,
	role models.Role,
) ([]models.Flat, error) {
	const op = `repo.House.GetFlatsByHouseId`

	builder := r.builder.
		Select(
			"flat_number",
			"house_id",
			"price",
			"rooms",
			"status",
		).
		From("flats").
		Where(sq.Eq{"house_id": houseId})

	if role == models.RoleClient {
		builder = builder.Where(sq.Eq{"status": models.StatusApproved})
	}

	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var flats []models.Flat
	for rows.Next() {
		var flat models.Flat
		if err := rows.Scan(
			&flat.Number,
			&flat.HouseID,
			&flat.Price,
			&flat.Rooms,
			&flat.Status,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		flats = append(flats, flat)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return flats, nil
}

func (r Repo) Update(
	ctx context.Context,
) error {
	const op = `repo.House.Update`
	builder := r.builder.
		Update("houses").
		Set("update_at", sq.Expr("CURRENT_TIMESTAMP"))

	sql, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r Repo) GetHouseById(
	ctx context.Context,
	houseId int,
) (models.House, error) {
	const op = `repo.Houses.GetHouseById`

	builder := r.builder.
		Select(
			"id",
			"address",
			"year",
			"developer",
			"created_at",
			"update_at",
		).
		From("houses").
		Where(sq.Eq{"id": houseId})

	sql, args, err := builder.ToSql()
	if err != nil {
		return models.House{}, fmt.Errorf("%s: %w", op, err)
	}

	var house models.House
	err = r.pool.QueryRow(ctx, sql, args...).Scan(
		&house.ID,
		&house.Address,
		&house.Year,
		&house.Developer,
		&house.CreatedAt,
		&house.UpdateAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.House{}, ErrHouseNotExists
		}
		return models.House{}, fmt.Errorf("%s: %w", op, err)
	}

	return house, nil
}
