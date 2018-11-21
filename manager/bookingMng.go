package manager

import (
	"fmt"
	"github.com/lian-rr/apartments-rental/repository"
	"time"
)

type Booking struct {
	ID        int
	Status    string
	StartDate time.Time
	EndDate   time.Time
	Details   string
	Apartment int
	Guest     int
	Active    bool
}

func ListBookings() (*[]*Booking, error) {

	repo, err := repository.BuildBookingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Booking Manager.", err)
	}

	defer repo.Close()

	bookings, err := repo.ListBookings()

	if err != nil {
		return nil, newError("Not possible to get list of bookings.", err)
	}

	return mapD2BookookList(bookings), nil
}

func FetchBooking(id int) (*Booking, error) {

	repo, err := repository.BuildBookingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Booking Manager.", err)
	}

	defer repo.Close()

	b, err := fetchBooking(id, repo)

	if err != nil {
		return nil, err
	}

	return mapD2Book(b), nil

}

func AddBooking(b *Booking) (*Booking, error) {

	repo, err := repository.BuildBookingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Booking Manager.", err)
	}

	defer repo.Close()

	nb, err := repo.PersistBooking(mapBook2D(b))

	if err != nil {
		return nil, newError("Not possible to persist new Booking.", err)
	}

	return mapD2Book(nb), nil
}

func UpdateBooking(b *Booking) (*Booking, error) {

	repo, err := repository.BuildBookingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Booking Manager.", err)
	}

	defer repo.Close()

	eB, err := fetchBooking(b.ID, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for Booking with id: %d", b.ID), err)
	}

	//Not found
	if eB == nil {
		return nil, nil
	}

	b.Active = eB.Active

	bg, err := repo.UpdateBooking(mapBook2D(b))

	if err != nil {
		return nil, newError("Not possible to update the Booking.", err)
	}

	return mapD2Book(bg), nil
}

func DeleteBooking(id int) (*Booking, error) {

	repo, err := repository.BuildBookingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Booking Manager.", err)
	}

	defer repo.Close()

	eB, err := fetchBooking(id, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for Booking with id: %d", id), err)
	}

	//Not found
	if eB == nil {
		return nil, nil
	}

	err = repo.DeleteBooking(id)

	if err != nil {
		return nil, newError("Not possible to delete the Booking.", err)
	}

	eB.Active = false

	return mapD2Book(eB), nil
}

func fetchBooking(id int, repo repository.BookingRepo) (*repository.Booking, error) {
	b, err := repo.FetchBooking(id)

	if err != nil {
		return nil, newError("Error fetching the booking information", err)
	}
	return b, nil
}

/*---------------------- ----- ---------------------*
/*---------------------- Utils ---------------------*
/*---------------------- ----- ---------------------*/

func mapBook2D(b *Booking) *repository.Booking {
	return &repository.Booking{
		ID:        b.ID,
		Status:    b.Status,
		StartDate: b.StartDate,
		EndDate:   b.EndDate,
		Details:   b.Details,
		Apartment: b.Apartment,
		Guest:     b.Guest,
		Active:    b.Active,
	}
}

func mapD2BookookList(bs *[]*repository.Booking) *[]*Booking {
	bl := make([]*Booking, 0)

	for _, b := range *bs {
		bl = append(bl, mapD2Book(b))
	}

	return &bl
}

func mapD2Book(b *repository.Booking) *Booking {
	if b != nil {
		return &Booking{
			ID:         b.ID,
			Status:    b.Status,
			StartDate: b.StartDate,
			EndDate:   b.EndDate,
			Details:   b.Details,
			Apartment: b.Apartment,
			Guest:     b.Guest,
			Active:    b.Active,
		}
	}
	return nil
}
