-- Migration: 010_create_listing_availability.sql
-- Description: Create listing_availability table for date-level availability and pricing

CREATE TABLE IF NOT EXISTS listing_availability (
    id SERIAL PRIMARY KEY,
    listing_id INT NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'available' CHECK (status IN ('available', 'booked', 'blocked')),
    price_override NUMERIC(15,2),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    UNIQUE (listing_id, date)
);

CREATE INDEX idx_listing_availability_listing_id ON listing_availability(listing_id);
CREATE INDEX idx_listing_availability_date ON listing_availability(date);
CREATE INDEX idx_listing_availability_status ON listing_availability(status);
CREATE INDEX idx_listing_availability_deleted_at ON listing_availability(deleted_at) WHERE deleted_at IS NULL;
