package entity

import (
	"encoding/json"
	"time"
)

type Property struct {
	ID                int64           `json:"id" db:"id"`
	OwnerID           int64           `json:"owner_id" db:"owner_id"`
	PropertyTypeID    int64           `json:"property_type_id" db:"property_type_id"`
	LocationID        int64           `json:"location_id" db:"location_id"`
	Address           string          `json:"address" db:"address"`
	Latitude          float64         `json:"latitude" db:"latitude"`
	Longitude         float64         `json:"longitude" db:"longitude"`
	AreaSqm           float64         `json:"area_sqm" db:"area_sqm"`
	BuildingAmenities json.RawMessage `json:"building_amenities,omitempty" db:"building_amenities"`
	CreatedAt         time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" db:"updated_at"`
	DeletedAt         *time.Time      `json:"deleted_at,omitempty" db:"deleted_at"`
}

type PropertyType struct {
	ID          int64      `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type PropertyWithDetails struct {
	Property
	PropertyTypeName string `json:"property_type_name,omitempty" db:"property_type_name"`
	LocationName     string `json:"location_name,omitempty" db:"location_name"`
	OwnerName        string `json:"owner_name,omitempty" db:"owner_name"`
}
