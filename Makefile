_COMPOSE=docker compose -f docker-compose.yaml --env-file .env -p avito-bootcamp-backend
_COMPOSE_TEST=docker compose -f tests/docker-compose-test.yaml --env-file tests/.env -p avito-bootcamp-backend-test

up:
	@echo "Starting Docker images..."
	${_COMPOSE} up -d

up_test:
	@echo "Starting Docker images..."
	${_COMPOSE_TEST} up -d
	
build:
	@echo "Building Docker images..."
	${_COMPOSE} build

build_test:
	@echo "Building Docker images..."
	${_COMPOSE_TEST} build

down:
	${_COMPOSE} down -v

down_test:
	${_COMPOSE_TEST} down -v

clean:
	${_COMPOSE} down --remove-orphans -v --rmi all

clean_test:
	${_COMPOSE_TEST} down --remove-orphans -v --rmi all

generate:
	go generate ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

cover-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

mockgen:
	mockgen -source=internal/service/houses/houses.go   -destination=internal/service/houses/houses_mocks_test.go -package=houses
	mockgen -source=internal/service/flats/flats.go   -destination=internal/service/flats/flats_mocks_test.go -package=flats