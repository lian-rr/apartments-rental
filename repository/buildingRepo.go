package repository

type (
	Building struct {
		ID          int
		SName       string
		FName       string
		Addr        string
		Phone       string
		Mng         string
		Description string
		Active      bool
	}

	BuildingRepo interface {
		ListBuildings() (*[]*Building, error)
		FetchBuilding(id int) (*Building, error)
		PersistBuilding(g *Building) (*Building, error)
		UpdateBuilding(g *Building) (*Building, error)
		DeleteBuilding(id int) error
		Close() error
	}
)
