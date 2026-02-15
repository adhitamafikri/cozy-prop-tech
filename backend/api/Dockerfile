# =============================================================================
# STAGE: builder - Compile the binary for production
# =============================================================================
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd

# =============================================================================
# TARGET: development - Run with Air for hot reload
# =============================================================================
FROM golang:1.26-alpine AS development

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8082

CMD ["air", "-c", "air.toml"]

# =============================================================================
# TARGET: production - Optimized image for deployment
# =============================================================================
FROM alpine AS production

WORKDIR /app

COPY --from=builder /bin/api /usr/local/bin/api

ENV GIN_MODE=release
EXPOSE 8082

CMD ["api"]
