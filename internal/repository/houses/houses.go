package houses

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
)

type Repo struct {
	pool    *pgxpool.Pool
	builder sq.StatementBuilderType
}

func New(pool *pgxpool.Pool) Repo {
	return Repo{
		pool:    pool,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r Repo) Create(
	ctx context.Context,
	dtoHouse dto.NewHouse,
) (models.House, error) {
	const op = `repo.House.Create`
	query := `
		INSERT INTO houses (address, year, developer)
		VALUES ($1, $2, $3)
		RETURNING id, address, year, developer, created_at, update_at
	`
	var house models.House
	err := r.pool.QueryRow(
		ctx,
		query,
		dtoHouse.Address,
		dtoHouse.Year,
		dtoHouse.Developer,
	).Scan(
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
		builder = builder.Where(sq.Eq{"status": "approved"}) // todo: create consts
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
