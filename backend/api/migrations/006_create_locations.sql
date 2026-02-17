-- Migration: 006_create_locations.sql
-- Description: Create locations table for hierarchical geographic data

CREATE TABLE IF NOT EXISTS locations (
    id SERIAL PRIMARY KEY,
    parent_id INT REFERENCES locations(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL CHECK (category IN ('COUNTRY', 'CITY', 'DISTRICT')),
    latitude NUMERIC(9,6) DEFAULT 0,
    longitude NUMERIC(9,6) DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_locations_parent_id ON locations(parent_id);
CREATE INDEX idx_locations_category ON locations(category);
CREATE INDEX idx_locations_deleted_at ON locations(deleted_at) WHERE deleted_at IS NULL;
