# syntax=docker/dockerfile:1

## Build
FROM golang AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . . 

RUN go build -o /checkinboard-server

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /checkinboard-server /checkinboard-server
COPY .env .env
COPY checkinboard-testing.db .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/checkinboard-server"]

