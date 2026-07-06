# Load .env if present (DATABASE_URL, PORT, etc.)
ifneq (,$(wildcard .env))
  include .env
  export
endif

DATABASE_URL ?= postgres://amirabaris@localhost:5432/goauth?sslmode=disable
MIGRATIONS_DIR := migrations

.PHONY: help sqlc migrate-up migrate-down migrate-status migrate-create dev run build

help: ## Show targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'

# --- sqlc ---
sqlc: ## Regenerate Go code from SQL queries
	sqlc generate

# --- goose ---
migrate-up: ## Apply all pending migrations
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

migrate-down: ## Roll back last migration
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

migrate-status: ## Show migration status
	goose -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status

migrate-create: ## Create migration: make migrate-create name=add_refresh_tokens
ifndef name
	$(error usage: make migrate-create name=your_migration_name)
endif
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# --- app ---
dev: ## Hot reload with air
	air

run: ## Run API once
	go run ./cmd/api

build: ## Build binary
	go build -o bin/api ./cmd/api

# --- combo targets ---
db: migrate-up sqlc ## Migrate up + regenerate sqlc

generate: sqlc ## Alias for sqlc