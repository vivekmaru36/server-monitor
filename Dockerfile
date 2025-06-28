# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o server-monitor .

# Stage 2: Lightweight image
FROM debian:bullseye-slim

WORKDIR /app

# Copy binary and static files
COPY --from=builder /app/server-monitor .
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server-monitor"]
