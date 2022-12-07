# syntax=docker/dockerfile:1

## Build
FROM golang AS build

WORKDIR /app

COPY . ./

RUN ls -la 
#RUN make build

##RUN go build -o /checkinboard-server
RUN make build 

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/checkinboard-server /checkinboard-server
COPY .env .env
COPY checkinboard-testing.db .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/checkinboard-server"]

