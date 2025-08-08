# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Final stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]