SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
CODE_COVERAGE=code-coverage

run: ## Start application
	$(GOCMD) run ./cmd/api

wire: ## Generate wire_gen.go
	cd pkg/di && wire

swag: ## Generate swagger docs
	swag init -g pkg/api/handler/offer.go -o ./cmd/api/docs

