# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# Final stage
FROM alpine:3.21
RUN apk --no-cache add ca-certificates tzdata

# Run as non-root user
RUN adduser -D -g '' appuser
USER appuser

WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/config ./config/

EXPOSE 8080
ENTRYPOINT ["/app/server"]

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/hello || exit 1
