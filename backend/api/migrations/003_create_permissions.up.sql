-- Migration: 003_create_permissions.sql
-- Description: Create permissions table for RBAC

CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_permissions_name ON permissions(name);
CREATE INDEX idx_permissions_deleted_at ON permissions(deleted_at) WHERE deleted_at IS NULL;
