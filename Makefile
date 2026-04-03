# Main Makefile for FitFeed project

.PHONY: help init dev dev-db dev-stop dev-auth dev-api dev-web migrate-up migrate-down

help:
	@echo "FitFeed Development Environment"
	@echo ""
	@echo "Usage:"
	@echo "  make init           - Check and install required tools"
	@echo "  make dev            - Run all services in development mode"
	@echo "  make dev-db         - Run only the database (Docker)"
	@echo "  make dev-stop       - Stop the database and other containers"
	@echo "  make dev-auth       - Run Auth service with hot-reload (Air)"
	@echo "  make dev-api        - Run API service with hot-reload (Air)"
	@echo "  make dev-web        - Run Web frontend (Vite)"
	@echo "  make migrate-up     - Run database migrations up"
	@echo "  make migrate-down   - Run database migrations down"

init:
	@echo "Checking development tools..."
	@command -v go >/dev/null 2>&1 || { echo >&2 "Go is not installed. Install it from https://golang.org/doc/install"; exit 1; }
	@command -v docker >/dev/null 2>&1 || { echo >&2 "Docker is not installed. Install it from https://docs.docker.com/get-docker/"; exit 1; }
	@command -v air >/dev/null 2>&1 || { echo >&2 "Air is not installed. Install it with: go install github.com/air-verse/air@latest"; exit 1; }
	@command -v node >/dev/null 2>&1 || { echo >&2 "Node.js is not installed. Install it from https://nodejs.org/"; exit 1; }
	@command -v bun >/dev/null 2>&1 || { echo >&2 "Bun is not installed. Install it with: curl -fsSL https://bun.sh/install | bash"; exit 1; }
	@if [ ! -f config.toml ]; then cp config.toml.template config.toml; echo "Created config.toml from template"; fi
	@echo "All tools found and config initialized."

dev-db:
	@echo "Starting Database..."
	@docker compose -f deployments/docker-compose/postgres/docker-compose.yml up -d

dev-stop:
	@echo "Stopping dev environment..."
	@docker compose -f deployments/docker-compose/postgres/docker-compose.yml down

dev-auth:
	@echo "Starting Auth service..."
	@FITFEED_CONF=$(PWD) air -c .air.auth.toml

dev-api:
	@echo "Starting API service..."
	@FITFEED_CONF=$(PWD) air -c .air.api.toml

dev-web:
	@echo "Starting Web frontend..."
	@cd services/web && bun install && bun run dev

migrate-up:
	@echo "Running migrations up..."
	@cd services/dbm && FITFEED_CONF=$(PWD) go run cmd/dbm/main.go up

migrate-down:
	@echo "Running migrations down..."
	@cd services/dbm && FITFEED_CONF=$(PWD) go run cmd/dbm/main.go down

dev: dev-db
	@echo "Waiting for database to be ready..."
	@sleep 3
	@make migrate-up
	@echo "Starting all services..."
	@make -j 3 dev-auth dev-api dev-web
