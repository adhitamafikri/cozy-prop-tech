-- Migration: 008_create_properties.sql
-- Description: Create properties table for physical real estate assets

CREATE TABLE IF NOT EXISTS properties (
    id SERIAL PRIMARY KEY,
    owner_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    property_type_id INT NOT NULL REFERENCES property_types(id) ON DELETE CASCADE,
    location_id INT NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    address TEXT NOT NULL,
    latitude NUMERIC(9,6) NOT NULL DEFAULT 0,
    longitude NUMERIC(9,6) NOT NULL DEFAULT 0,
    area_sqm DOUBLE PRECISION DEFAULT 0,
    building_amenities JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_properties_owner_id ON properties(owner_id);
CREATE INDEX idx_properties_property_type_id ON properties(property_type_id);
CREATE INDEX idx_properties_location_id ON properties(location_id);
CREATE INDEX idx_properties_deleted_at ON properties(deleted_at) WHERE deleted_at IS NULL;
