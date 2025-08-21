# Name app
APP_NAME = server
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= root:root@tcp(127.0.0.1:3306)/shopdevgo
GOOSE_MIGRATION_DIR ?= sql/schema

ifeq ($(OS),Windows_NT)
	SET_ENV = set
else
	SET_ENV = export
endif

docker_build:
	docker-compose up -d --build
	docker-compose ps
docker_stop:
	docker-compose down
docker_up:
	docker compose up -d

up_by_one:
	@$(SET_ENV) GOOSE_DRIVER=$(GOOSE_DRIVER)&&$(SET_ENV) GOOSE_DBSTRING=$(GOOSE_DBSTRING)&&goose -dir=$(GOOSE_MIGRATION_DIR) up-by-one

dev:
	go run ./cmd/$(APP_NAME)

# create new migration
create_migration:
	goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql

upse:
	@$(SET_ENV) GOOSE_DRIVER=$(GOOSE_DRIVER)&&$(SET_ENV) GOOSE_DBSTRING=$(GOOSE_DBSTRING)&&goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@$(SET_ENV) GOOSE_DRIVER=$(GOOSE_DRIVER)&&$(SET_ENV) GOOSE_DBSTRING=$(GOOSE_DBSTRING)&&goose -dir=$(GOOSE_MIGRATION_DIR) down

resetse:
	@$(SET_ENV) GOOSE_DRIVER=$(GOOSE_DRIVER)&&$(SET_ENV) GOOSE_DBSTRING=$(GOOSE_DBSTRING)&&goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate

swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: dev downse upse resetse docker_build docker_stop docker_up create_migration sqlgen

.PHONY: air
