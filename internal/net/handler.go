package net

import (
	authapi "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/auth"
	flatsapi "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/flats"
	housesapi "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/houses"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/auth"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/flats"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/houses"
	"go.uber.org/zap"
)

type (
	houseAPI = housesapi.API
	flatAPI  = flatsapi.API
	authAPI  = authapi.API
)

type Handler struct {
	houseAPI
	flatAPI
	authAPI
}

func NewHandler(authService auth.Service, houseService houses.Service, flatService flats.Service, logger *zap.Logger) Handler {
	return Handler{
		houseAPI: housesapi.New(houseService, logger),
		flatAPI:  flatsapi.New(flatService, logger),
		authAPI:  authapi.New(authService, logger),
	}
}
