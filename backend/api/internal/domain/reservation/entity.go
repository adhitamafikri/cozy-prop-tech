package entity

import "time"

type Reservation struct {
	ID         int64     `json:"id" db:"id"`
	ListingID  int64     `json:"listing_id" db:"listing_id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	CheckIn    time.Time `json:"check_in" db:"check_in"`
	CheckOut   time.Time `json:"check_out" db:"check_out"`
	TotalPrice float64   `json:"total_price" db:"total_price"`
	NumGuests  int       `json:"num_guests" db:"num_guests"`
	Status     string    `json:"status" db:"status"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type ReservationWithDetails struct {
	Reservation
	ListingTitle string `json:"listing_title,omitempty" db:"listing_title"`
	UserName     string `json:"user_name,omitempty" db:"user_name"`
	UserEmail    string `json:"user_email,omitempty" db:"user_email"`
}
