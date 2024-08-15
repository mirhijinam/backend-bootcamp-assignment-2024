package flats

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models/dto"
)

// newFlatNumberQuery - запрос для получения нового номера квартиры с учетом последней добавленной в дом.
// Если в доме нет квартир, задается значение по умолчанию 1
const newFlatNumberQuery = `
	(
		SELECT COALESCE(MAX(flat_number) + 1, 1)
		FROM flats
		WHERE house_id = ?
	)`

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
	draftFlat dto.NewFlat,
) (models.Flat, error) {
	const op = `repo.Flats.Create`

	builder := r.builder.
		Insert("flats").
		Columns(
			"flat_number",
			"house_id",
			"price",
			"rooms",
		).
		Values(
			sq.Expr(newFlatNumberQuery, draftFlat.HouseID),
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

	fmt.Println(sql)

	var flat models.Flat
	err = r.pool.QueryRow(ctx, sql, args...).Scan(
		&flat.Number,
		&flat.HouseID,
		&flat.Price,
		&flat.Rooms,
		&flat.Status,
	)
	if err != nil {
		return models.Flat{}, fmt.Errorf("%s: %w", op, err)
	}

	return flat, nil
}
