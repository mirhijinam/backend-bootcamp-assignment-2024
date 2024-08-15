package auth

import "go.uber.org/zap"

type service interface {
}

type API struct {
	logger *zap.Logger
}

func New(service service, logger *zap.Logger) API {
	return API{
		logger: logger,
	}
}
