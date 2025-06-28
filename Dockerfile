# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o monitor .

# Stage 2: Lightweight image
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/monitor .

EXPOSE 8080

# CMD ["./monitor"]
CMD ["./server-monitor"]
