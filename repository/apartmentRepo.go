package repository

type (
	Apartment struct {
		ID       int
		Number   string
		Baths    float32
		Beds     int
		Rooms    int
		Building int
		Details  string
		Active   bool
	}

	ApartmentRepo interface {
		ListApartments() (*[]*Apartment, error)
		FetchApartment(id int) (*Apartment, error)
		PersistApartment(g *Apartment) (*Apartment, error)
		UpdateApartment(g *Apartment) (*Apartment, error)
		DeleteApartment(id int) error
		Close() error
	}
)
