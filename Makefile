_COMPOSE=docker compose -f docker-compose.yaml --env-file .env -p avito-bootcamp-backend

up:
	@echo "Starting Docker images..."
	${_COMPOSE} up

build:
	@echo "Building Docker images..."
	${_COMPOSE} build

down:
	${_COMPOSE} down

clean:
	${_COMPOSE} down --remove-orphans -v --rmi all

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