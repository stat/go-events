GO					= go
GO_BUILD		= $(GO) build
GO_CLEAN		= $(GO) clean
GO_TEST			= $(GO) test
GO_GET			= $(GO) get

SRC					:= $(shell find . -type f -name '*.go' -not -path "./vendor/*")

TARGET			= main
BUILDS			= ./build
TESTS				= ./...

BUILD				:= `git rev-parse HEAD`
MAJOR				:= $(shell cat .version | cut -d. -f1)
MINOR				:= $(shell cat .version | cut -d. -f2)
PATCH				:= $(shell cat .version | cut -d. -f3)

GOFLAGS			= -v
LDFLAGS			= -X=$(TARGET)/pkg/version.Build=$(BUILD) \
							-X=$(TARGET)/pkg/version.Major=$(MAJOR) \
							-X=$(TARGET)/pkg/version.Minor=$(MINOR) \
							-X=$(TARGET)/pkg/version.Patch=$(PATCH) \
							-X=$(TARGET)/pkg/version.String=${MAJOR}.${MINOR}.${PATCH}

RELEASEFLAGS := GOOS=linux GOARCH=arm64 CGO_ENABLED=0

.DEFAULT_GOAL := help

ifneq (,$(wildcard ./.env))
	include .env
endif

export

$(BUILDS)/$(TARGET): $(SRC)
	@$(GO_BUILD) $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(BUILDS)/$(TARGET) .
	@echo Binary "(v${MAJOR}.${MINOR}.${PATCH})" built to $(BUILDS)/$(TARGET)!

.PHONY: clean
clean: ## remove the build directory
	@$(RM) -r $(BUILDS)

.PHONY: test
test: ## run tests
	@$(GO_TEST) -ldflags="$(LDFLAGS)" -v $(TESTS)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
