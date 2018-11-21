package manager

import (
	"fmt"
	"github.com/lian-rr/apartments-rental/repository"
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

func FetchGuest(id int) (*Guest, error) {

	repo, err := repository.BuildGuestRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Guest Manager.", err)
	}

	defer repo.Close()

	g, err := fetchGuest(id, repo)

	if err != nil {
		return nil, err
	}

	return mapD2G(g), nil

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

func UpdateGuest(g *Guest) (*Guest, error) {

	repo, err := repository.BuildGuestRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Guest Manager.", err)
	}

	defer repo.Close()

	eG, err := fetchGuest(g.ID, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for guest with id: %d", g.ID), err)
	}

	//Not found
	if eG == nil {
		return nil, nil
	}

	g.Active = eG.Active

	ng, err := repo.UpdateGuest(mapG2D(g))

	if err != nil {
		return nil, newError("Not possible to update the guest.", err)
	}

	return mapD2G(ng), nil

}

func DeleteGuest(id int) (*Guest, error) {

	repo, err := repository.BuildGuestRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Guest Manager.", err)
	}

	defer repo.Close()

	eG, err := fetchGuest(id, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for guest with id: %d", id), err)
	}

	//Not found
	if eG == nil {
		return nil, nil
	}

	err = repo.DeleteGuest(id)

	if err != nil {
		return nil, newError("Not possible to delete the guest.", err)
	}

	eG.Active = false

	return mapD2G(eG), nil
}

func fetchGuest(id int, repo repository.GuestRepo) (*repository.Guest, error) {
	g, err := repo.FetchGuest(id)

	if err != nil {
		return nil, newError("Error fetching the guest information", err)
	}
	return g, nil
}

/*---------------------- ----- ---------------------*
/*---------------------- Utils ---------------------*
/*---------------------- ----- ---------------------*/


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
