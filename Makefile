GO					= go
GO_BUILD		= $(GO) build
GO_CLEAN		= $(GO) clean
GO_TEST			= $(GO) clean -testcache; $(GO) test
GO_GET			= $(GO) get

SRC					:= $(shell find . -path "./\.*" -prune -type f -name '*.go')

TARGET			= main
BUILDS			= ./build
TESTS				= ./pkg/...

BUILD				:= `git rev-parse HEAD`
MAJOR				:= $(shell cat .version | cut -d. -f1)
MINOR				:= $(shell cat .version | cut -d. -f2)
PATCH				:= $(shell cat .version | cut -d. -f3)

GOFLAGS			= -v
GOTAGS 			= viper_bind_struct
LDFLAGS			= -X=$(TARGET)/pkg/version.Build=$(BUILD) \
							-X=$(TARGET)/pkg/version.Major=$(MAJOR) \
							-X=$(TARGET)/pkg/version.Minor=$(MINOR) \
							-X=$(TARGET)/pkg/version.Patch=$(PATCH) \
							-X=$(TARGET)/pkg/version.String=${MAJOR}.${MINOR}.${PATCH}

RELEASEFLAGS := GOOS=linux GOARCH=amd64 CGO_ENABLED=0

CURRENT_UID := $(id -u)

.DEFAULT_GOAL := help

ifneq (,$(wildcard ./.env))
	include .env
endif

ifdef ENV
ifneq (,$(wildcard ./.env-${ENV}))
	include .env-${ENV}
endif
endif

ifndef CASSANDRA_CONNECTION_STRING
	CASSANDRA_CONNECTION_STRING="cassandra://${CASSANDRA_HOSTS}:${CASSANDRA_CQL_PORT}/${CASSANDRA_KEYSPACE}"
endif

ifndef POSTGRES_CONNECTION_STRING
	POSTGRES_CONNECTION_STRING="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"
endif

export

$(BUILDS)/$(TARGET): $(SRC)
	@$(GO_BUILD) $(GOFLAGS) -ldflags="$(LDFLAGS)" -tags="$(GOTAGS)" -o $(BUILDS)/$(TARGET) .
	@echo Binary "(v${MAJOR}.${MINOR}.${PATCH})" built to $(BUILDS)/$(TARGET)!

.PHONY: benchmark
benchmark: ## run benchmarks
	@$(GO_TEST) -ldflags="$(LDFLAGS)" -tags="${GOTAGS}" -v -bench=. $(TESTS)

.PHONY: build 
build: $(BUILDS)/$(TARGET) ## build for development

.PHONY: clean
clean: ## remove the build directory
	@$(RM) -r $(BUILDS)

.PHONY: compose-up
compose-up: release
	@docker-compose up --build --remove-orphans

.PHONY: compose-down
compose-down:
	@docker-compose down

.PHONY: migrate-install
migrate-install:
	@go install -tags 'cassandra postgres file' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: cass-migrate-create
cass-migrate-create: migrate-install
	@migrate create -ext cql -dir ./migrations/cassandra -seq $(n)

.PHONY: cass-migrate-down
cass-migrate-down: migrate-install
	@migrate -database="${CASSANDRA_CONNECTION_STRING}" -path=./migrations/cassandra -lock-timeout=30 -verbose down 1

.PHONY: cass-migrate-force
cass-migrate-force: migrate-install
	@migrate -database="${CASSANDRA_CONNECTION_STRING}" -path=./migrations/cassandra -lock-timeout=30 -verbose force ${v}

.PHONY: cass-migrate-up
cass-migrate-up: migrate-install
	@migrate -database="${CASSANDRA_CONNECTION_STRING}" -path=./migrations/cassandra -lock-timeout=30 -verbose up

.PHONY: pg-migrate-create
pg-migrate-create: migrate-install
	@migrate create -ext sql -dir ./migrations/postgres -seq $(n)

.PHONY: pg-migrate-down
pg-migrate-down: migrate-install
	@migrate -database="${POSTGRES_CONNECTION_STRING}" -path=./migrations/postgres -lock-timeout=30 -verbose down 1

.PHONY: pg-migrate-force
pg-migrate-force: migrate-install
	@migrate -database="${POSTGRES_CONNECTION_STRING}" -path=./migrations/postgres -lock-timeout=30 -verbose force ${v}

.PHONY: pg-migrate-up
pg-migrate-up: migrate-install
	@migrate -database="${POSTGRES_CONNECTION_STRING}" -path=./migrations/postgres -lock-timeout=30 -verbose up

release:
	@$(RELEASEFLAGS) $(GO_BUILD) $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(BUILDS)/release .

.PHONY: run
run: clean build ## run with .env
	@$(BUILDS)/$(TARGET)

.PHONY: test
test: ## run tests
	# @$(GO_TEST) -ldflags="$(LDFLAGS)" -tags="${GOTAGS}" -v $(TESTS)
	@$(GO_TEST) -ldflags="$(LDFLAGS)" -tags="${GOTAGS}" -v -run TestPipeline ./pkg/tasks

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
