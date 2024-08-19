package tests

import (
	"context"
	"net/http/httptest"
	"testing"

	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/lamoda/gonkey/fixtures"
	"github.com/lamoda/gonkey/runner"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/generated"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/config"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/net/middleware"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/pkg/db"
	authrepo "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/auth"
	flatsrepo "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/flats"
	houserepo "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/repository/houses"
	authservice "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/auth"
	flatsservice "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/flats"
	housesservice "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/service/houses"
	"go.uber.org/zap"
)

func TestFuncCases(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	pool, err := db.MustOpenDB(context.Background(), cfg.DBConfig)
	if err != nil {
		panic(err)
	}

	lgr, err := zap.NewDevelopment()
	if err != nil {
		t.Fatal(err)
	}

	transactor := manager.Must(trmpgx.NewDefaultFactory(pool))
	authService := authservice.New(
		authrepo.New(pool), lgr)

	houseService := housesservice.New(
		houserepo.New(pool), lgr)

	flatsService := flatsservice.New(
		flatsrepo.New(pool), houserepo.New(pool), lgr, transactor)

	authMiddl := middleware.New(lgr, cfg.ServerConfig.SecretKey)

	server, err := generated.NewServer(
		net.NewHandler(authService, houseService, flatsService, lgr, cfg.ServerConfig.SecretKey),
		authMiddl,
	)
	if err != nil {
		panic(err)
	}

	// запустите выполнение тестов из директории cases с записью в отчет Allure
	runner.RunWithTesting(t, &runner.RunWithTestingParams{
		Server:      httptest.NewServer(server),
		TestsDir:    "cases",
		DB:          stdlib.OpenDBFromPool(pool),
		DbType:      fixtures.Postgres,
		FixturesDir: "fixtures",
	})
}
