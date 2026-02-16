# ERD Analysis - Cozy Prop Tech

**Date**: 2026-02-16
**Status**: Initial Review
**ERD Location**: `docs/system-design/erd.png`

---

## Executive Summary

The current ERD provides a solid foundation for the Cozy Prop Tech platform, covering the core domains (Auth, Users, Properties, Listings, Availability, Booking). However, several critical enhancements are needed for production readiness, particularly around RBAC scalability, payment processing, audit trails, and geolocation support.

---

## Current Schema Overview

### Identified Tables
1. **users** - Authentication and user management
2. **user_profiles** - Extended user profile information
3. **properties** - Physical property assets
4. **listings** - Rentable units/rooms within properties
5. **amenities** - Master list of property amenities
6. **property_amenities** - Junction table (properties ↔ amenities)
7. **availability** - Listing availability calendar
8. **bookings** - Reservation records
9. **payments** - Payment transaction records

---

## ✅ Strengths

### 1. Domain Separation
- Clear separation of concerns across required domains
- Auth, Users, Properties, Listings, Availability, Booking are all represented

### 2. Data Normalization
- Good separation between `users` and `user_profiles`
- Smart distinction between `properties` (physical assets) and `listings` (rentable units)
- Proper many-to-many relationship via `property_amenities`

### 3. Standard Practices
- Timestamp fields (`created_at`, `updated_at`) across tables
- Foreign key relationships properly established
- Status fields for state management

---

## 🔍 Critical Issues & Recommendations

### 1. RBAC Implementation

**Current State:**
- ✅ `users.role` field exists
- ⚠️ Appears to be enum-based

**Issues:**
- Limited scalability for complex permission requirements
- Hard to manage granular permissions (e.g., "can approve bookings", "can view reports")
- Admin panel will need fine-grained access control

**Recommendation:**
```sql
-- Enhanced RBAC schema
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    resource VARCHAR(50) NOT NULL, -- e.g., 'bookings', 'properties'
    action VARCHAR(50) NOT NULL,   -- e.g., 'create', 'read', 'update', 'delete'
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE role_permissions (
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- Update users table
ALTER TABLE users ADD COLUMN role_id UUID REFERENCES roles(id);
```

**Decision Point:** Keep simple enum for now and migrate later, or implement full RBAC from start?

---

### 2. Property Ownership

**Current State:**
- ⚠️ No clear owner relationship visible in `properties` table

**Issues:**
- Cannot identify who owns/manages each property
- Multi-tenant property management not supported
- Cannot filter properties by landlord/manager

**Recommendation:**
```sql
ALTER TABLE properties
ADD COLUMN owner_id UUID REFERENCES users(id) NOT NULL,
ADD COLUMN manager_id UUID REFERENCES users(id); -- Optional property manager
```

**Use Cases:**
- Landlords viewing their portfolio
- Property managers handling multiple properties
- Admin panel filtering by owner
- Revenue reporting per landlord

---

### 3. Booking Status Management

**Current State:**
- ✅ `bookings.status` field exists

**Ensure Complete Lifecycle:**
```sql
CREATE TYPE booking_status AS ENUM (
    'pending',      -- Initial booking request
    'confirmed',    -- Payment received, booking confirmed
    'checked_in',   -- Guest has arrived
    'checked_out',  -- Stay completed
    'cancelled',    -- Cancelled before check-in
    'no_show',      -- Guest didn't arrive
    'refunded'      -- Payment refunded
);
```

**Additional Fields Needed:**
- `cancellation_reason` TEXT
- `cancelled_at` TIMESTAMP
- `cancelled_by` UUID (user_id)

---

### 4. Availability Model Enhancement

**Current State:**
- `availability` table tracks listing availability

**Questions to Address:**
1. **Date Range Blocking**: How to block dates for maintenance?
2. **Recurring Patterns**: How to define "available every weekend"?
3. **Minimum Stay**: Different min nights for different periods?
4. **Dynamic Pricing**: Seasonal rates, weekend premiums?

**Recommended Schema:**
```sql
CREATE TABLE availability (
    id UUID PRIMARY KEY,
    listing_id UUID REFERENCES listings(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    status VARCHAR(20) NOT NULL, -- 'available', 'booked', 'blocked'
    price DECIMAL(10,2),         -- Daily price (can override base price)
    minimum_stay INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(listing_id, date)
);

CREATE INDEX idx_availability_listing_date ON availability(listing_id, date);
```

**Alternative Approach (More Efficient):**
```sql
-- Block/unblock date ranges instead of daily records
CREATE TABLE availability_blocks (
    id UUID PRIMARY KEY,
    listing_id UUID REFERENCES listings(id) ON DELETE CASCADE,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL,
    reason TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    CHECK (end_date >= start_date)
);
```

---

### 5. Payment System Enhancements

**Current State:**
- ✅ `payments` table exists with `booking_id` FK

**Critical Missing Fields:**
```sql
ALTER TABLE payments ADD COLUMN
    payment_method VARCHAR(50),           -- 'credit_card', 'bank_transfer', 'paypal'
    payment_provider VARCHAR(50),         -- 'stripe', 'paypal', 'manual'
    transaction_id VARCHAR(255),          -- External payment gateway reference
    currency CHAR(3) DEFAULT 'USD',       -- ISO 4217 currency code
    amount_paid DECIMAL(10,2) NOT NULL,
    service_fee DECIMAL(10,2),            -- Platform fee
    tax_amount DECIMAL(10,2),
    total_amount DECIMAL(10,2) NOT NULL,
    refund_amount DECIMAL(10,2) DEFAULT 0,
    refund_status VARCHAR(20),            -- 'none', 'partial', 'full'
    refunded_at TIMESTAMP,
    metadata JSONB;                       -- Store provider-specific data
```

**Payment States:**
```sql
CREATE TYPE payment_status AS ENUM (
    'pending',
    'processing',
    'completed',
    'failed',
    'cancelled',
    'refunded',
    'partially_refunded'
);
```

---

### 6. Soft Deletes

**Current State:**
- ⚠️ No `deleted_at` fields visible

**Issues:**
- Hard deletes lose audit trail
- Cannot recover accidentally deleted data
- Breaks foreign key integrity for historical records

**Recommendation:**
Add to ALL tables:
```sql
ALTER TABLE users ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE properties ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE listings ADD COLUMN deleted_at TIMESTAMP;
ALTER TABLE bookings ADD COLUMN deleted_at TIMESTAMP;
-- etc.
```

**Query Pattern:**
```sql
-- All queries should filter out soft-deleted records
SELECT * FROM users WHERE deleted_at IS NULL;
```

**sqlc Integration:**
```sql
-- name: GetActiveUsers :many
SELECT * FROM users WHERE deleted_at IS NULL;

-- name: SoftDeleteUser :exec
UPDATE users SET deleted_at = NOW() WHERE id = $1;
```

---

### 7. Listing Pricing

**Critical Missing Fields:**
```sql
ALTER TABLE listings ADD COLUMN
    base_price DECIMAL(10,2) NOT NULL,
    cleaning_fee DECIMAL(10,2) DEFAULT 0,
    currency CHAR(3) DEFAULT 'USD',
    pricing_type VARCHAR(20) NOT NULL, -- 'per_night', 'per_month', 'per_week'
    deposit_amount DECIMAL(10,2) DEFAULT 0,
    min_stay_nights INT DEFAULT 1,
    max_stay_nights INT,
    instant_book BOOLEAN DEFAULT FALSE;
```

---

### 8. Geolocation for Maps

**Required for Leaflet + OpenStreetMap Integration:**
```sql
ALTER TABLE properties ADD COLUMN
    latitude DECIMAL(10,8) NOT NULL,
    longitude DECIMAL(11,8) NOT NULL,
    address_line1 VARCHAR(255) NOT NULL,
    address_line2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    state_province VARCHAR(100),
    postal_code VARCHAR(20),
    country CHAR(2) NOT NULL; -- ISO 3166-1 alpha-2

-- For efficient geospatial queries
CREATE INDEX idx_properties_location ON properties USING gist(
    point(longitude, latitude)
);
```

**Consider PostGIS Extension:**
```sql
CREATE EXTENSION postgis;

ALTER TABLE properties ADD COLUMN
    location GEOGRAPHY(POINT, 4326);

-- Update location from lat/long
UPDATE properties SET location = ST_SetSRID(ST_MakePoint(longitude, latitude), 4326);

-- Find properties within 10km of a point
SELECT * FROM properties
WHERE ST_DWithin(
    location,
    ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
    10000 -- meters
);
```

---

### 9. Audit Trail

**Recommendation for Compliance:**
```sql
CREATE TABLE booking_history (
    id UUID PRIMARY KEY,
    booking_id UUID REFERENCES bookings(id) ON DELETE CASCADE,
    status_from VARCHAR(20),
    status_to VARCHAR(20) NOT NULL,
    changed_by UUID REFERENCES users(id),
    changed_at TIMESTAMP DEFAULT NOW(),
    notes TEXT
);

CREATE TABLE payment_logs (
    id UUID PRIMARY KEY,
    payment_id UUID REFERENCES payments(id) ON DELETE CASCADE,
    event_type VARCHAR(50) NOT NULL, -- 'created', 'captured', 'refunded'
    metadata JSONB,
    created_at TIMESTAMP DEFAULT NOW()
);
```

---

### 10. Database Indexing

**Essential Indexes:**
```sql
-- Users
CREATE UNIQUE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);

-- Bookings
CREATE INDEX idx_bookings_listing_id ON bookings(listing_id);
CREATE INDEX idx_bookings_user_id ON bookings(user_id);
CREATE INDEX idx_bookings_status ON bookings(status);
CREATE INDEX idx_bookings_dates ON bookings(check_in_date, check_out_date);
CREATE INDEX idx_bookings_created_at ON bookings(created_at DESC);

-- Listings
CREATE INDEX idx_listings_property_id ON listings(property_id);
CREATE INDEX idx_listings_status ON listings(status); -- if exists

-- Availability
CREATE INDEX idx_availability_listing_date ON availability(listing_id, date);

-- Payments
CREATE INDEX idx_payments_booking_id ON payments(booking_id);
CREATE INDEX idx_payments_status ON payments(status);
CREATE INDEX idx_payments_created_at ON payments(created_at DESC);

-- Foreign Keys (if not auto-indexed)
-- PostgreSQL automatically creates indexes on primary keys and unique constraints
-- But verify FKs are indexed for JOIN performance
```

---

## 📋 Missing Tables to Consider

### 1. Reviews/Ratings System

**Critical for Property Platforms:**
```sql
CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    booking_id UUID REFERENCES bookings(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    listing_id UUID REFERENCES listings(id) ON DELETE CASCADE,
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    cleanliness_rating INT CHECK (cleanliness_rating BETWEEN 1 AND 5),
    accuracy_rating INT CHECK (accuracy_rating BETWEEN 1 AND 5),
    communication_rating INT CHECK (communication_rating BETWEEN 1 AND 5),
    location_rating INT CHECK (location_rating BETWEEN 1 AND 5),
    value_rating INT CHECK (value_rating BETWEEN 1 AND 5),
    comment TEXT,
    response TEXT,            -- Host response to review
    responded_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(booking_id)       -- One review per booking
);

CREATE INDEX idx_reviews_listing ON reviews(listing_id);
CREATE INDEX idx_reviews_user ON reviews(user_id);
```

---

### 2. Notification System

**For User Engagement:**
```sql
CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,     -- 'booking_confirmed', 'payment_received', etc.
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    link VARCHAR(500),             -- Deep link to relevant page
    read_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_notifications_user ON notifications(user_id, created_at DESC);
CREATE INDEX idx_notifications_unread ON notifications(user_id, read_at) WHERE read_at IS NULL;
```

---

### 3. Favorites/Wishlists

**Improve User Experience:**
```sql
CREATE TABLE user_favorites (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    listing_id UUID REFERENCES listings(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (user_id, listing_id)
);

CREATE INDEX idx_favorites_user ON user_favorites(user_id, created_at DESC);
```

---

### 4. Property Images

**If Not Handled via External Service:**
```sql
CREATE TABLE property_images (
    id UUID PRIMARY KEY,
    property_id UUID REFERENCES properties(id) ON DELETE CASCADE,
    listing_id UUID REFERENCES listings(id) ON DELETE CASCADE, -- Nullable, some images for property, some for specific listings
    url VARCHAR(500) NOT NULL,
    thumbnail_url VARCHAR(500),
    is_primary BOOLEAN DEFAULT FALSE,
    display_order INT DEFAULT 0,
    caption TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_images_property ON property_images(property_id, display_order);
CREATE INDEX idx_images_listing ON property_images(listing_id, display_order);
```

---

### 5. Email Verification & Password Reset

**For Secure Auth Flow:**
```sql
CREATE TABLE email_verifications (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE password_resets (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) UNIQUE NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_email_verifications_token ON email_verifications(token);
CREATE INDEX idx_password_resets_token ON password_resets(token);
```

---

## 🎯 PostgreSQL-Specific Recommendations

### 1. Use UUIDs for Primary Keys

**Benefits:**
- No ID enumeration attacks
- Better for distributed systems
- Merge-friendly (no ID conflicts)

**Implementation:**
```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    -- other fields
);
```

---

### 2. JSONB for Flexible Data

**Use Cases:**
```sql
ALTER TABLE properties ADD COLUMN
    features JSONB; -- { "pool": true, "garden": true, "parking_spots": 2 }

ALTER TABLE listings ADD COLUMN
    house_rules JSONB; -- { "no_smoking": true, "no_pets": false, "quiet_hours": "22:00-08:00" }

ALTER TABLE users ADD COLUMN
    preferences JSONB; -- User settings, notification preferences, etc.

-- Query JSONB
SELECT * FROM properties WHERE features->>'pool' = 'true';
SELECT * FROM properties WHERE features @> '{"pool": true}';

-- Create GIN index for JSONB queries
CREATE INDEX idx_properties_features ON properties USING gin(features);
```

---

### 3. ENUM Types for Type Safety

**Benefits:**
- Database-level validation
- Better performance than VARCHAR
- Self-documenting

**Implementation:**
```sql
CREATE TYPE user_role AS ENUM ('guest', 'host', 'admin', 'super_admin');
CREATE TYPE booking_status AS ENUM ('pending', 'confirmed', 'checked_in', 'checked_out', 'cancelled', 'refunded');
CREATE TYPE payment_status AS ENUM ('pending', 'processing', 'completed', 'failed', 'cancelled', 'refunded');
CREATE TYPE listing_status AS ENUM ('draft', 'active', 'inactive', 'archived');

ALTER TABLE users ADD COLUMN role user_role DEFAULT 'guest';
ALTER TABLE bookings ADD COLUMN status booking_status DEFAULT 'pending';
ALTER TABLE payments ADD COLUMN status payment_status DEFAULT 'pending';
ALTER TABLE listings ADD COLUMN status listing_status DEFAULT 'draft';
```

---

### 4. CHECK Constraints

**Data Integrity at Database Level:**
```sql
ALTER TABLE bookings ADD CONSTRAINT check_dates
    CHECK (check_out_date > check_in_date);

ALTER TABLE bookings ADD CONSTRAINT check_future_booking
    CHECK (check_in_date >= CURRENT_DATE);

ALTER TABLE payments ADD CONSTRAINT check_positive_amount
    CHECK (amount_paid >= 0);

ALTER TABLE reviews ADD CONSTRAINT check_rating_range
    CHECK (rating BETWEEN 1 AND 5);

ALTER TABLE listings ADD CONSTRAINT check_capacity
    CHECK (max_guests > 0);
```

---

### 5. Triggers for Automated Fields

**Auto-update `updated_at`:**
```sql
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_properties_updated_at BEFORE UPDATE ON properties
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Apply to all tables with updated_at
```

---

### 6. Partitioning for Large Tables

**For High-Volume Tables (bookings, availability):**
```sql
-- Partition bookings by year
CREATE TABLE bookings (
    id UUID NOT NULL,
    -- other fields
    check_in_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
) PARTITION BY RANGE (check_in_date);

CREATE TABLE bookings_2025 PARTITION OF bookings
    FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');

CREATE TABLE bookings_2026 PARTITION OF bookings
    FOR VALUES FROM ('2026-01-01') TO ('2027-01-01');
```

---

## 🚀 Next Steps & Action Items

### Phase 1: Critical Fixes (Week 1)
- [ ] Add `owner_id` to properties table
- [ ] Enhance payments table with provider/transaction fields
- [ ] Add geolocation fields (latitude, longitude) to properties
- [ ] Add pricing fields to listings (base_price, currency, etc.)
- [ ] Implement soft deletes across all tables

### Phase 2: Feature Completeness (Week 2)
- [ ] Decide on RBAC approach (simple enum vs. full RBAC tables)
- [ ] Implement reviews/ratings system
- [ ] Add property images table
- [ ] Enhance availability model (decide on daily records vs. blocks)
- [ ] Add email verification & password reset tables

### Phase 3: Optimization (Week 3)
- [ ] Create all necessary indexes
- [ ] Implement ENUM types for status fields
- [ ] Add CHECK constraints
- [ ] Set up triggers for `updated_at`
- [ ] Add JSONB fields for flexible data

### Phase 4: Advanced Features (Week 4)
- [ ] Implement notification system
- [ ] Add favorites/wishlists
- [ ] Set up audit trail tables
- [ ] Consider PostGIS for advanced geospatial queries
- [ ] Evaluate partitioning strategy for high-volume tables

---

## 🤔 Questions for Clarification

1. **Multi-currency Support**: Do we need to support multiple currencies from day 1?
2. **RBAC Complexity**: How granular should permissions be? (simple roles vs. resource-action permissions)
3. **Pricing Strategy**: Fixed pricing vs. dynamic pricing? Seasonal rates?
4. **Availability Model**: Daily records or date range blocks? Both?
5. **Image Storage**: Database vs. object storage (S3/Cloudinary)?
6. **Booking Conflicts**: How to handle overlapping booking attempts?
7. **Cancellation Policy**: Flexible, moderate, strict? Automated refund calculation?
8. **Payment Gateway**: Stripe, PayPal, both? Manual payments allowed?

---

## 📚 References

- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [sqlc Best Practices](https://docs.sqlc.dev/en/latest/howto/index.html)
- [PostGIS for Geospatial](https://postgis.net/)
- [Database Normalization](https://en.wikipedia.org/wiki/Database_normalization)
- [RBAC Design Patterns](https://en.wikipedia.org/wiki/Role-based_access_control)

---

## Conclusion

The current ERD provides a **solid foundation** but requires **significant enhancements** for production readiness. Focus on:

1. **Payment integrity** (critical for business operations)
2. **Geolocation support** (required for map features)
3. **Proper pricing model** (core business logic)
4. **Property ownership** (multi-tenant support)
5. **Audit trails** (compliance & debugging)

Recommended approach: **Iterative implementation** - start with Phase 1 critical fixes, validate with real use cases, then expand features in subsequent phases.
