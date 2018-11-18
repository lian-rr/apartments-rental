package repository

import "time"

type (
	Guest struct {
		ID      int
		Fname   string
		Lname   string
		Bdate   time.Time
		Gender  string
		Details string
		Active  bool
	}

	GuestRepo interface {
		ListGuests() (*[]*Guest, error)
		FetchGuest(id int) (*Guest, error)
		PersistGuest(g *Guest) (*Guest, error)
		UpdateGuest(g *Guest) (*Guest, error)
		DeleteGuest(id int) error
		Close() error
	}
)
