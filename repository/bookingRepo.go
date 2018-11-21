package repository

import "time"

type (
	Booking struct {
		ID        int
		Status    string
		StartDate time.Time
		EndDate   time.Time
		Details   string
		Apartment int
		Guest     int
		Active    bool
	}

	BookingRepo interface {
		ListBookings() (*[]*Booking, error)
		FetchBooking(id int) (*Booking, error)
		PersistBooking(g *Booking) (*Booking, error)
		UpdateBooking(g *Booking) (*Booking, error)
		DeleteBooking(id int) error
		Close() error
	}
)
