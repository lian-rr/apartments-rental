package manager

import (
	"fmt"
	"time"
	"github.com/lian-rr/apartment-rental/repository"
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

func AddGuest(g *Guest) (*Guest, error) {

	gRepo, err := repository.BuildGuestRepo()

	if err != nil {
		fmt.Printf("Not posible to initiate the Guest Manager")
		return &Guest{}, err
	}

	ng, err := gRepo.PersistGuest(mapG2D(g))

	if err != nil {
		fmt.Printf("Not posible to persist new Gues")
		return &Guest{}, err
	}

	return mapD2G(ng), nil

}

func mapG2D(g *Guest) *repository.Guest {
	return &repository.Guest{ID: g.ID, Fname: g.Fname, Lname: g.Lname, Bdate: g.Bdate, Gender: g.Gender, Details: g.Details, Active: g.Active}
}

func mapD2G(g *repository.Guest) *Guest {
	return &Guest{ID: g.ID, Fname: g.Fname, Lname: g.Lname, Bdate: g.Bdate, Gender: g.Gender, Details: g.Details, Active: g.Active}
}
