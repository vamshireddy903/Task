# Build Stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o go-app

# Run Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/go-app .
EXPOSE 8080
ENTRYPOINT ["./go-app"]
