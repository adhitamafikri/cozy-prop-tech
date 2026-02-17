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

# Defaults (fallback)
DB_USER="${DB_USER:-cozy}"
DB_PASSWORD="${DB_PASSWORD:-cozy123}"
DB_NAME="${DB_NAME:-cozy_prop_db}"
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"

# Override to localhost for psql (host machine connection)
DB_HOST="localhost"

exec psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME"
