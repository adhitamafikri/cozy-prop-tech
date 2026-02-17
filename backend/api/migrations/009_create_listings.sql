-- Migration: 009_create_listings.sql
-- Description: Create listings table for rental units

CREATE TABLE IF NOT EXISTS listings (
    id SERIAL PRIMARY KEY,
    property_id INT NOT NULL REFERENCES properties(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    base_price NUMERIC(15,2),
    status VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'inactive')),
    guest_capacity INT NOT NULL DEFAULT 0,
    minimum_reservation_nights INT NOT NULL,
    maximum_reservation_nights INT NOT NULL,
    cleaning_fee NUMERIC(15,2),
    extra_guest_capacity INT DEFAULT 0,
    extra_guest_fee NUMERIC(15,2),
    num_beds INT NOT NULL,
    num_bathrooms INT NOT NULL,
    num_bedrooms INT NOT NULL,
    amenities JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_listings_property_id ON listings(property_id);
CREATE INDEX idx_listings_status ON listings(status);
CREATE INDEX idx_listings_deleted_at ON listings(deleted_at) WHERE deleted_at IS NULL;
