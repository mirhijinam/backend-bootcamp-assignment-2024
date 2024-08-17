package models

import "github.com/google/uuid"

type Status string

const (
	StatusCreated      Status = "created"
	StatusOnModeration Status = "on_moderation"
	StatusApproved     Status = "approved"
	StatusDeclined     Status = "declined"
)

type Flat struct {
	Number      int
	HouseID     int
	Price       int
	Rooms       int
	Status      Status
	ModeratorID *uuid.UUID
}
