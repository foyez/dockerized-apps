include app.env
export $(shell sed 's/=.*//' app.env)

## up_dev: up docker containers for development
up_dev:
	@echo "upping docker containers for development..."
	docker compose --env-file app.env -f docker-compose.yaml -f docker-compose.dev.yaml up --build

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

## create_migration: create migration up & down files
create_migration:
	@echo "creating migration files..."
	@mkdir -p db
	migrate create -ext sql -dir db/migration -seq ${name}

## migrateup: apply all up migrations
migrateup:
	@echo "applying all up migrations..."
	migrate -path db/migration -database ${DB_SOURCE} -verbose up

## migratedown: apply all down migrations
migratedown:
	@echo "applying all down migrations..."
	migrate -path db/migration -database ${DB_SOURCE} -verbose down

## sqlc: generate Go code from SQL
sqlc:
	@echo "generating go code from sql..."
	sqlc generate

## test: test the project
test:
	@echo "testing the project..."
	go test -v -cover ./...

.PHONY: up_dev up down_dev down create_migration migrateup migratedown sqlc test