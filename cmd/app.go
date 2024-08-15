package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/config"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/middleware"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/pkg/db"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/pkg/logger"
	flatsrepo "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/flats"
	houserepo "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/houses"
	flatsservice "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/flats"
	houseservice "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/houses"
	"go.uber.org/zap"
)

func Run() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	pool, err := db.MustOpenDB(context.Background(), cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	houseService := houseservice.New(
		houserepo.New(pool),
	)

	flatsService := flatsservice.New(
		flatsrepo.New(pool),
	)

	lgr := logger.New(cfg.LoggerConfig.Mode, cfg.LoggerConfig.Filepath)
	defer lgr.Sync()

	authMiddl := middleware.New(lgr)

	server, err := generated.NewServer(
		net.NewHandler(houseService, flatsService, lgr),
		authMiddl,
	)
	if err != nil {
		panic(err)
	}

	lgr.Info("Server is running", zap.String("port", cfg.ServerConfig.Port))
	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		server,
	); err != nil {
		panic(err)
	}
}
