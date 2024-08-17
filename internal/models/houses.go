package models

import "time"

type House struct {
	ID        int
	Address   string
	Year      int
	Developer *string
	CreatedAt time.Time
	UpdateAt  *time.Time
}
