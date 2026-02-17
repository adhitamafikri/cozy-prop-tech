package entity

import "time"

type Location struct {
	ID        int64      `json:"id" db:"id"`
	ParentID  *int64     `json:"parent_id,omitempty" db:"parent_id"`
	Name      string     `json:"name" db:"name"`
	Category  string     `json:"category" db:"category"`
	Latitude  float64    `json:"latitude" db:"latitude"`
	Longitude float64    `json:"longitude" db:"longitude"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type LocationWithChildren struct {
	Location
	Children []Location `json:"children,omitempty"`
}
