FROM golang:alpine as builder
LABEL org.opencontainers.image.authors="michzuerch@gmail.com"
RUN mkdir /build 
WORKDIR /build
COPY / ./
#RUN ls -la
RUN go mod tidy
RUN go mod vendor 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
#FROM scratch
FROM ubuntu
COPY --from=builder /build/main /app/
COPY .env /app/
RUN ls -la
WORKDIR /app
CMD ["./main"]
