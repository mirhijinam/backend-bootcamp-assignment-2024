package dto

import (
	"github.com/google/uuid"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"
)

type NewFlat struct {
	Number  int
	HouseID int
	Price   int
	Rooms   int
}

type FlatUpdateParams struct {
	Number      int
	HouseID     int
	Status      models.Status
	ModeratorID uuid.UUID
}
