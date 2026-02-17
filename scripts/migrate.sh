#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"

# Load environment from .env
if [ -f "$PROJECT_DIR/backend/api/.env" ]; then
    set -a
    source "$PROJECT_DIR/backend/api/.env"
    set +a
fi

# Defaults (fallback if .env not found)
DB_USER="${DB_USER:-cozy}"
DB_PASSWORD="${DB_PASSWORD:-cozy123}"
DB_NAME="${DB_NAME:-cozy_prop_db}"
DB_HOST="${DB_HOST:-postgres}"
DB_PORT="${DB_PORT:-5432}"
DB_SSLMODE="${DB_SSLMODE:-disable}"
DOCKER_NETWORK="${DOCKER_NETWORK:-cozy-prop-tech_cozy-network}"

DIRECTION="${1:-up}"
STEPS="${2:-1}"

DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"
MIGRATIONS_PATH="$PROJECT_DIR/backend/api/migrations"

case "$DIRECTION" in
    up)
        docker run --rm --network "$DOCKER_NETWORK" -v "$MIGRATIONS_PATH:/migrations" \
            migrate/migrate \
            -path=/migrations \
            -database="$DATABASE_URL" \
            up
        ;;
    down)
        docker run --rm --network "$DOCKER_NETWORK" -v "$MIGRATIONS_PATH:/migrations" \
            migrate/migrate \
            -path=/migrations \
            -database="$DATABASE_URL" \
            down "$STEPS"
        ;;
    *)
        echo "Usage: $0 [up|down] [steps]"
        exit 1
        ;;
esac
