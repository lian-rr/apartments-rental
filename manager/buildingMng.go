package manager

import (
	"fmt"
	"github.com/lian-rr/apartment-rental/repository"
)

type Building struct {
	ID          int
	SName       string
	FName       string
	Addr        string
	Phone       string
	Mng         string
	Description string
	Active      bool
}

func ListBuildings() (*[]*Building, error) {

	repo, err := repository.BuildBuildingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Building Manager.", err)
	}

	defer repo.Close()

	buildings, err := repo.ListBuildings()

	if err != nil {
		return nil, newError("Not possible to get list of buildings.", err)
	}

	return mapD2BList(buildings), nil
}

func FetchBuilding(id int) (*Building, error) {

	repo, err := repository.BuildBuildingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Building Manager.", err)
	}

	defer repo.Close()

	b, err := fetchBuilding(id, repo)

	if err != nil {
		return nil, err
	}

	return mapD2B(b), nil

}

func AddBuilding(g *Building) (*Building, error) {

	repo, err := repository.BuildBuildingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Building Manager.", err)
	}

	defer repo.Close()

	nb, err := repo.PersistBuilding(mapB2D(g))

	if err != nil {
		return nil, newError("Not possible to persist new Building.", err)
	}

	return mapD2B(nb), nil
}

func UpdateBuilding(b *Building) (*Building, error) {

	repo, err := repository.BuildBuildingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Building Manager.", err)
	}

	defer repo.Close()

	eB, err := fetchBuilding(b.ID, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for Building with id: %d", b.ID), err)
	}

	//Not found
	if eB == nil {
		return nil, nil
	}

	b.Active = eB.Active

	bg, err := repo.UpdateBuilding(mapB2D(b))

	if err != nil {
		return nil, newError("Not possible to update the Building.", err)
	}

	return mapD2B(bg), nil
}

func DeleteBuilding(id int) (*Building, error) {

	repo, err := repository.BuildBuildingRepo()

	if err != nil {
		return nil, newError("Not possible to initiate the Building Manager.", err)
	}

	defer repo.Close()

	eB, err := fetchBuilding(id, repo)

	if err != nil {
		return nil, newError(fmt.Sprintf("Error retrieving the data for Building with id: %d", id), err)
	}

	//Not found
	if eB == nil {
		return nil, nil
	}

	err = repo.DeleteBuilding(id)

	if err != nil {
		return nil, newError("Not possible to delete the Building.", err)
	}

	eB.Active = false

	return mapD2B(eB), nil
}


func fetchBuilding(id int, repo repository.BuildingRepo) (*repository.Building, error) {
	b, err := repo.FetchBuilding(id)

	if err != nil {
		return nil, newError("Error fetching the building information", err)
	}
	return b, nil
}

/*---------------------- ----- ---------------------*
/*---------------------- Utils ---------------------*
/*---------------------- ----- ---------------------*/

func mapB2D(b *Building) *repository.Building {
	return &repository.Building{
		ID:          b.ID,
		SName:       b.SName,
		FName:       b.FName,
		Addr:        b.Addr,
		Phone:       b.Phone,
		Mng:         b.Mng,
		Description: b.Description,
		Active:      b.Active,
	}
}

func mapD2BList(bs *[]*repository.Building) *[]*Building {
	bl := make([]*Building, 0)

	for _, b := range *bs {
		bl = append(bl, mapD2B(b))
	}

	return &bl
}

func mapD2B(b *repository.Building) *Building {
	if b != nil {
		return &Building{
			ID:          b.ID,
			SName:       b.SName,
			FName:       b.FName,
			Addr:        b.Addr,
			Phone:       b.Phone,
			Mng:         b.Mng,
			Description: b.Description,
			Active:      b.Active,
		}
	}
	return nil
}
