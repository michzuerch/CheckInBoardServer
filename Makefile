ifneq (,$(wildcard ./.env))
    include .env
    export
endif

build:
	@echo ${HTTP_PORT}
	@echo "Start building..."
	go mod tidy
#	go generate
	go build main.go
	@echo  "Building done."

build-docker:
	@echo "Start build docker image..."
	docker build --tag checkinboard-server .
	docker image tag checkinboard-server:latest checkinboard-server:v1.0
	@echo "Docker images is ready."

run:
	go run main.go

run-docker: clean build-docker
	docker run --publish ${HTTP_PORT}:${HTTP_PORT} --name checkinboard-server checkinboard-server

test-models:
	go test ./models


clean:
	docker stop checkinboard-server
	docker container rm checkinboard-server

all: build

