package manager

import (
	"fmt"
	"github.com/lian-rr/apartment-rental/repository"
)

type Apartment struct {
	ID       int
	Number   string
	Baths    float32
	Beds     int
	Rooms    int
	Building int
	Details  string
	Active   bool
}

func ListApartments() (*[]*Apartment, error) {

	repo, err := repository.BuildApartmentRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Apartment Manager.", err)
	}

	defer repo.Close()

	apartments, err := repo.ListApartments()

	if err != nil {
		return nil, newError("Not possible to get list of apartments.", err)
	}

	return mapA2DList(apartments), nil
}

func FetchApartment(id int) (*Apartment, error) {

	repo, err := repository.BuildApartmentRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Apartment Manager.", err)
	}

	defer repo.Close()

	a, err := fetchApartment(id, repo)

	if err != nil {
		return nil, err
	}

	return mapD2A(a), nil

}

func AddApartment(a *Apartment) (*Apartment, error) {

	repo, err := repository.BuildApartmentRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Apartment Manager.", err)
	}

	defer repo.Close()

	na, err := repo.PersistApartment(mapA2D(a))

	if err != nil {
		return nil, newError("Not possible to persist new Apartment.", err)
	}

	return mapD2A(na), nil
}

func UpdateApartment(a *Apartment) (*Apartment, error) {

	repo, err := repository.BuildApartmentRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Apartment Manager.", err)
	}

	defer repo.Close()

	aB, err := fetchApartment(a.ID, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for Apartment with id: %d", a.ID), err)
	}

	//Not found
	if aB == nil {
		return nil, nil
	}

	a.Active = aB.Active

	ag, err := repo.UpdateApartment(mapA2D(a))

	if err != nil {
		return nil, newError("Not possible to update the Apartment.", err)
	}

	return mapD2A(ag), nil
}

func DeleteApartment(id int) (*Apartment, error) {

	repo, err := repository.BuildApartmentRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Apartment Manager.", err)
	}

	defer repo.Close()

	aB, err := fetchApartment(id, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for Apartment with id: %d", id), err)
	}

	//Not found
	if aB == nil {
		return nil, nil
	}

	err = repo.DeleteApartment(id)

	if err != nil {
		return nil, newError("Not possible to delete the Apartment.", err)
	}

	aB.Active = false

	return mapD2A(aB), nil
}

func fetchApartment(id int, repo repository.ApartmentRepo) (*repository.Apartment, error) {
	a, err := repo.FetchApartment(id)

	if err != nil {
		return nil, newError("Error fetching the apartment information", err)
	}
	return a, nil
}

/*---------------------- ----- ---------------------*
/*---------------------- Utils ---------------------*
/*---------------------- ----- ---------------------*/

func mapA2D(b *Apartment) *repository.Apartment {
	return &repository.Apartment{
		ID:       b.ID,
		Number:   b.Number,
		Baths:    b.Baths,
		Beds:     b.Beds,
		Rooms:    b.Rooms,
		Details:  b.Details,
		Building: b.Building,
		Active:   b.Active,
	}
}

func mapA2DList(bs *[]*repository.Apartment) *[]*Apartment {
	al := make([]*Apartment, 0)

	for _, a := range *bs {
		al = append(al, mapD2A(a))
	}

	return &al
}

func mapD2A(b *repository.Apartment) *Apartment {
	if b != nil {
		return &Apartment{
			ID:       b.ID,
			Number:   b.Number,
			Baths:    b.Baths,
			Beds:     b.Beds,
			Rooms:    b.Rooms,
			Details:  b.Details,
			Building: b.Building,
			Active:   b.Active,
		}
	}
	return nil
}
