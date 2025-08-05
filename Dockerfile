# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-app

# Run stage
FROM scratch
COPY --from=builder /app/go-app /go-app
EXPOSE 5000
ENTRYPOINT ["/go-app"]
