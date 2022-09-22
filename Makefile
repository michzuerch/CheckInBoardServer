.DEFAULT_GOAL:=help
SHELL:=/bin/zsh

PROJECT_NAME := "CheckInBoardServer"
BINARY_NAME := main

# sqlite3, mysql, psql, mssql
# SQL_MIGRATE_DB := "sqlite3"

PKG := "github.com/michzuerch/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
BUILD_DATE := $(shell date +'%Y-%m-%d')

.PHONY: all dep lint vet test test-coverage build clean

ifneq (,$(wildcard ./.env))
    include .env
    export
endif



help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	
##@ Golang

all: build ## Build all
	$(info Build)

dep: ## Downloag go dependencies
	$(info Checking and getting dependencies)
	go mod download

lint: ## Linting
	$(info Linting with golangci-list)
	golangci-lint run --enable-all

install-golangci-lint: ## Install golangci-lint
	$(info Install golangci-lint)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0

vet: ## go vet
	$(info Report likely misstaken packages)
	go vet ${PKG_LIST}

build: ## go mod download, go build
	$(info go mod download, go build)
	go get -u
	go mod tidy
	go mod vendor
	go build -o ${BINARY_NAME}

run: ## Run the application on local machine
	$(info Run main on local machine)
	./${BINARY_NAME}

release: ## goreleaser
	$(info goreleaser)
	goreleaser release --snapshot --rm-dist

##@ Testing

test: ## Testing
	$(info Unit tests, go test)
	go test -short ${PKG_LIST}

test-coverage: ## Test coverage
	$(info Coverage)
	go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}

##@ Docker

docker-build: ## Build the docker image
	$(info Build the docker image)
	@echo "Start build docker image..."
	docker build --no-cache --tag checkinboard-server --build-arg buildDate=$(BUILD_DATE) .
	docker image tag checkinboard-server:latest checkinboard-server:v1.0
	@echo "Docker images is ready."

docker-run: ## Run the application as docker container
	$(info Start a docker container with main.go)
	@echo "Docker container on port:" ${HTTP_PORT}
	docker run --publish ${HTTP_PORT}:${HTTP_PORT} --name checkinboard-server checkinboard-server

docker-clean: ## Stop the running docker container and remove the container
	$(info docker stop, docker container rm)
	docker stop checkinboard-server
	docker container rm checkinboard-server

docker-shell: ## Connet to shell inside docker container
	$(info Shell inside docker container)
	docker run -it --name checkinboard-server checkinboard-server bash

##@ Database

database-postgres-start: ## Start postgres with docker-compose
	$(info Start the postgres database)
	docker-compose -p postgres -f ./Database/docker-compose-Postgres.yml up -d

database-postgres-stop: ## Stop postgres with docker-compose
	$(info Stop the postgres database)
	docker-compose -p postgres -f ./Database/docker-compose-Postgres.yml down

database-mysql-start: ## Start mysql with docker-compose
	$(info Start the mysql database)
	docker-compose -p mysql -f ./Database/docker-compose-Mysql.yml up -d

database-mysql-stop: ## Stop mysql with docker-compose
	$(info Stop the mysql database)
	docker-compose -p mysql -f ./Database/docker-compose-Mysql.yml down

database-mssql-start: ## Start mssql with docker-compose
	$(info Start the mssql database)
	docker-compose -p mssql -f ./Database/docker-compose-Mssqlserver.yml up -d

database-mssql-stop: ## Stop mssql with docker-compose
	$(info Stop the mssql database)
	docker-compose -p mssql -f ./Database/docker-compose-Mssqlserver.yml down

##@ sql-migrate

install-sql-migrate: ## Install sql-migrate
	$(info Install sql-migrate)
	go install github.com/rubenv/sql-migrate/...@latest

sql-migrate-up: ## sql-migrate up
	$(info sql-migrate up)
	sql-migrate up -env=${DB_TYPE}

sql-migrate-down: ## sql-migrate down
	$(info sql-migrate down)
	sql-migrate down -env=${DB_TYPE}

sqlboiler: ## sqlboiler
	$(info sqlboiler)
	sqlboiler ${DB_TYPE}

database-test-sqlite: ## Test the migration status
	$(info Test the migration status)
	sqlite3 ${DB_CONNECT_STRING} "SELECT COUNT(*) FROM migrations"

##@ Cleanup

clean: ## Cleanup the project folders
	$(info Cleaning up folders)
	go clean
	#rm ${BINARY_NAME}

