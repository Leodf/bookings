PROJECT_NAME = bookings
MODULE_NAME = cmd

.SILENT:
.DEFAULT_GOAL := help

.PHONY: help
help:
	$(info bookings commands:)
# $(info -> setup                   installs dependencies)
# $(info -> format                  formats go files)
# $(info -> build                   builds executable binary)
	$(info -> test                    runs available tests)
# $(info -> run                     runs application)
	$(info -> coverage                generate test coverage report)

# .PHONY: setup
# setup:
# 	go get -d -v -t ./...
# 	go install -v ./...
# 	go mod tidy -v

# .PHONY: format
# format:
# 	go fmt ./...

# .PHONY: build
# build:
# 	go build -v -o $(MODULE_NAME).bin ./$(MODULE_NAME)
# 	chmod +x $(MODULE_NAME).bin
# 	echo $(MODULE_NAME).bin

.PHONY: test
test:
	go test ./... -v -covermode=count

# .PHONY: run
# run:
# 	go run ./$(MODULE_NAME)

.PHONY: coverage
coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o cover.html
