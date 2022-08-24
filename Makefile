build:
	@echo "Start building..."
	go mod tidy
	go generate
	go build main.go
	@echo  "Building done."

build-docker:
	@echo "Start build docker image..."
	docker build --tag checkinboard-server .
	docker image tag checkinboard-server:latest checkinboard-server:v1.0
	@echo "Docker images is ready."

run:
	go run main.go

run-docker:
	docker run --publish 8080:8080 --name checkinboard-server checkinboard-server

test-models:
	go test ./models

all: build

