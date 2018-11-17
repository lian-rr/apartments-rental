package manager

import (
	"fmt"
	"github.com/lian-rr/apartment-rental/repository"
	"time"
)

type Guest struct {
	ID      int
	Fname   string
	Lname   string
	Bdate   time.Time
	Gender  string
	Details string
	Active  bool
}

func ListGuests() (*[]*Guest, error) {

	repo, err := repository.BuildGuestRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Guest Manager.", err)
	}

	defer repo.Close()

	guests, err := repo.ListGuests()

	if err != nil {
		return nil, newError("Not possible to get list of guests.", err)
	}

	return mapD2GList(guests), nil
}

func AddGuest(g *Guest) (*Guest, error) {

	repo, err := repository.BuildGuestRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Guest Manager.", err)
	}

	defer repo.Close()

	ng, err := repo.PersistGuest(mapG2D(g))

	if err != nil {
		return nil, newError("Not possible to persist new Guest.", err)
	}

	return mapD2G(ng), nil
}

func FetchGuest(id int) (*Guest, error) {

	repo, err := repository.BuildGuestRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Guest Manager.", err)
	}

	defer repo.Close()

	g, err := repo.FetchGuest(id)

	if err != nil {
		return nil, newError("Error fetching the guest information", err)
	}

	return mapD2G(g), nil

}

/*---------------------- ----- ---------------------*
/*---------------------- Utils ---------------------*
/*---------------------- ----- ---------------------*/

func newError(m string, err error) error {
	fmt.Printf("Error: %s", err.Error())
	return fmt.Errorf("%s\n", m)
}

func mapG2D(g *Guest) *repository.Guest {
	return &repository.Guest{
		ID:      g.ID,
		Fname:   g.Fname,
		Lname:   g.Lname,
		Bdate:   g.Bdate,
		Gender:  g.Gender,
		Details: g.Details,
		Active:  g.Active,
	}
}

func mapD2GList(gs *[]*repository.Guest) *[]*Guest {
	gl := make([]*Guest, 0)

	for _, g := range *gs {
		gl = append(gl, mapD2G(g))
	}

	return &gl
}

func mapD2G(g *repository.Guest) *Guest {
	if g != nil {
		return &Guest{
			ID:      g.ID,
			Fname:   g.Fname,
			Lname:   g.Lname,
			Bdate:   g.Bdate,
			Gender:  g.Gender,
			Details: g.Details,
			Active:  g.Active,
		}
	}
	return nil
}
