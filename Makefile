GOPATH=$(shell go env GOPATH)
BINARY_NAME=realestate
BUILD_HASH=$(shell git rev-parse HEAD)
MIGRATION_NAME ?= new_migration
MIGRATION_DATE ?= `date '+%Y%m%d%H%M%S' -u`
MIGRATION_PATH ?= ./internal/infrastructure/datastore/migrations

export GOPROXY=http://proxy.golang.org,direct
export GOSUMDB=off

.PHONE: run
run: build
	go run ./cmd/estate

.PHONE: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./dist/$(BINARY_NAME) -x \
 		 -ldflags="-X main.build=$(BUILD_HASH) -s -w -extldflags '-static'" \
 		  ./cmd/estate

.PHONY: docker-build
docker-build:
	docker build --build-arg BINARY_NAME=$(BINARY_NAME) --build-arg BUILD_HASH=$(BUILD_HASH) -t realestate/api .

.PHONY: up
up: docker-build
	docker compose up

.PHONY: upd
upd: docker-build
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: lint
lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2
	${GOPATH}/bin/golangci-lint run --timeout 10m --fix --skip-dirs tests --config .golangci.yml

.PHONY: gen-mocks
gen-mocks:
	go install github.com/vektra/mockery/v2@v2.20.0
	mockery --case snake --dir ./internal/app --all

.PHONY: test
test:
	GOOS=linux go test -cover `go list ./internal/...`

.PHONY: fmt
fmt:
	gofmt -w -s cmd
	gofmt -w -s pkg

.PHONY: gen-swag
gen-swag:
	swag init -g cmd/estate/main.go -o docs/swagger --parseDepth 10 --parseDependency --parseInternal \
		--outputTypes go,json --parseGoList false --propertyStrategy camelcase --pd true -d ./internal/adapters/handlers, ./internal

.PHONY: migration
migration:
	@DATE= && \
	touch "$(MIGRATION_PATH)/$(MIGRATION_DATE)_$(name).tx.up.sql" && \
	touch "$(MIGRATION_PATH)/$(MIGRATION_DATE)_$(name).tx.down.sql"


.PHONY: help
help:
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "up", "run the application"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "upd", "run the application in background"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "down", "shutdown the application"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "lint", "run linter"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "fmt", "run gofmt"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "gen-mocks", "generate mocks"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "test", "run go test"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "build", "run go build"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "docker-build", "create docker image"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "run", "run go run"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "help", "print this help"}'
	@echo "" | awk '{printf "\033[36m%-30s\033[0m %s\n", "migration", "creates new migration file with name and date, use name=your_migration_name"}'