FROM golang:alpine as builder
RUN mkdir /build 
WORKDIR /build
COPY / ./
COPY .env ./
RUN ls -la
#RUN go get -u
#RUN go mod tidy
RUN go mod vendor 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
#RUN ls -la
#FROM scratch
FROM ubuntu
ARG buildDate
LABEL org.opencontainers.image.authors="michzuerch@gmail.com"
LABEL org.opencontainers.image.created=$buildDate
LABEL org.opencontainers.image.url="https://github.com/michzuerch/CheckInBoard-Server"
LABEL org.opencontainers.image.version="0.1.0"
LABEL org.opencontainers.image.title="CheckInBoard Server"
COPY --from=builder /build/main /app/
COPY .env /app/
RUN ls -la /app
WORKDIR /app
CMD ["./main"]
