
FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache redis

WORKDIR /app


COPY --from=builder /app/server ./
COPY --from=builder /app/templates ./templates/
COPY .env /app

EXPOSE 8080


CMD ["./server"]
