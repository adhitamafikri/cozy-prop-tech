#!/bin/bash

if [ -f "$(dirname "$0")/../.env" ]; then
    set -a
    source "$(dirname "$0")/../.env"
    set +a
fi

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_NAME="${DB_NAME:-cozy_prop_db}"
DB_USER="${DB_USER:-cozy}"
DB_PASSWORD="${DB_PASSWORD:-cozy123}"

if [ -z "$1" ]; then
    docker compose exec -it postgres psql -U "$DB_USER" -d "$DB_NAME"
else
    docker compose exec -it postgres psql -U "$DB_USER" -d "$DB_NAME" -c "$1"
fi
