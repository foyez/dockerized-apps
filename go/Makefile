include app.env
export $(shell sed 's/=.*//' app.env)

## up_dev: up docker containers for development
up_dev:
	@echo "upping docker containers for development..."
	docker compose --env-file app.env -f docker-compose.yaml -f docker-compose.dev.yaml up -d --build

## up: up docker containers
up:
	@echo "upping docker containers..."
	docker compose --env-file app.env up -d --build

## down_dev: remove docker containers for development
down_dev:
	@echo "removing docker containers for development..."
	docker compose --env-file app.env -f docker-compose.yaml -f docker-compose.dev.yaml down

## down: remove docker containers
down:
	@echo "removing docker containers..."
	docker compose --env-file app.env down

## test: test the project
test:
	@echo "testing the project..."
	go test -v -cover ./...

.PHONY: up_dev up down_dev down test