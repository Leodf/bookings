include .env
PROJECT_NAME = bookings
MODULE_NAME = cmd

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info bookings commands:)
# $(info -> setup                   installs dependencies)
# $(info -> format                  formats go files)
	$(info -> build                   builds executable binary)
	$(info -> test                    runs available tests)
	$(info -> run                     runs application)
	$(info -> coverage                generate test coverage report)
	$(info -> create_migration        generate new migration files, param "NAME=" it's necessary)
	$(info -> db_status               dump the migration status for the current DB)
	$(info -> migrations_up           migrate the DB to the most recent version available)
	$(info -> migrations_down         roll back the version by 1)
	$(info -> migrations_reset        roll back all migrations)

# .PHONY: setup
# setup:
# 	go get -d -v -t ./...
# 	go install -v ./...
# 	go mod tidy -v

# .PHONY: format
# format:
# 	go fmt ./...

.PHONY: build
build:
	go build -v -o ./build/bookings cmd/web/*.go
# chmod +x $(MODULE_NAME).bin
# echo $(MODULE_NAME).bin

.PHONY: test
test:
	go test ./... -covermode=count

.PHONY: run
run:
	go build -v -o ./build/bookings cmd/web/*.go && ./build/bookings

.PHONY: coverage
coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o cover.html

.PHONY: create_migration
create_migration:
	goose -dir=$(GOOSE_MIGRATION_DIR) create $(NAME) sql

.PHONY: db_status
db_status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DSN) goose -dir=$(GOOSE_MIGRATION_DIR) status

.PHONY: migrations_up
migrations_up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DSN) goose -dir=$(GOOSE_MIGRATION_DIR) up

.PHONY: migrations_down
migrations_down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DSN) goose -dir=$(GOOSE_MIGRATION_DIR) down

.PHONY: migrations_reset
migrations_reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DSN) goose -dir=$(GOOSE_MIGRATION_DIR) reset
