# Development stage
FROM golang:1.22.8-alpine AS development

LABEL authors="n1207n@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY .. .

CMD ["go", "run", "cmd/app/main.go"]

# Test stage
FROM development AS test

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux

CMD ["go", "test", "./..."]

# Build stage
FROM development AS build

WORKDIR /app

# Build the binary
RUN go build -o main cmd/main.go

# Production stage
FROM alpine AS production

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]
