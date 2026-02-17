package listing

import (
	"encoding/json"
	"time"
)

type Listing struct {
	ID                       int64           `json:"id" db:"id"`
	PropertyID               int64           `json:"property_id" db:"property_id"`
	Title                    string          `json:"title" db:"title"`
	Description              string          `json:"description,omitempty" db:"description"`
	BasePrice                float64         `json:"base_price" db:"base_price"`
	Status                   string          `json:"status" db:"status"`
	GuestCapacity            int             `json:"guest_capacity" db:"guest_capacity"`
	MinimumReservationNights int             `json:"minimum_reservation_nights" db:"minimum_reservation_nights"`
	MaximumReservationNights int             `json:"maximum_reservation_nights" db:"maximum_reservation_nights"`
	CleaningFee              float64         `json:"cleaning_fee" db:"cleaning_fee"`
	ExtraGuestCapacity       int             `json:"extra_guest_capacity" db:"extra_guest_capacity"`
	ExtraGuestFee            float64         `json:"extra_guest_fee" db:"extra_guest_fee"`
	NumBeds                  int             `json:"num_beds" db:"num_beds"`
	NumBathrooms             int             `json:"num_bathrooms" db:"num_bathrooms"`
	NumBedrooms              int             `json:"num_bedrooms" db:"num_bedrooms"`
	Amenities                json.RawMessage `json:"amenities,omitempty" db:"amenities"`
	CreatedAt                time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt                time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt                *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}

type ListingAvailability struct {
	ID            int64      `json:"id" db:"id"`
	ListingID     int64      `json:"listing_id" db:"listing_id"`
	Date          time.Time  `json:"date" db:"date"`
	Status        string     `json:"status" db:"status"`
	PriceOverride *float64   `json:"price_override,omitempty" db:"price_override"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type ListingImage struct {
	ID        int64     `json:"id" db:"id"`
	ListingID int64     `json:"listing_id" db:"listing_id"`
	URL       string    `json:"url" db:"url"`
	IsPrimary bool      `json:"is_primary" db:"is_primary"`
	Order     int       `json:"order" db:"order"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ListingWithDetails struct {
	Listing
	PropertyTitle   string `json:"property_title,omitempty" db:"property_title"`
	PropertyAddress string `json:"property_address,omitempty" db:"property_address"`
}
