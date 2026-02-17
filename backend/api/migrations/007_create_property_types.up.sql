-- Migration: 007_create_property_types.sql
-- Description: Create property_types table for property classification

CREATE TABLE IF NOT EXISTS property_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_property_types_name ON property_types(name);
CREATE INDEX idx_property_types_deleted_at ON property_types(deleted_at) WHERE deleted_at IS NULL;
