package repository

import (
	"database/sql"
	"fmt"
)

type ApartmentRepoSQL struct {
	db *sql.DB
}

func newApartmentRepoSQL(db *sql.DB) ApartmentRepoSQL {
	return ApartmentRepoSQL{db: db}
}

func (m ApartmentRepoSQL) ListApartments() (*[]*Apartment, error) {
	rows, err := m.db.Query(`SELECT id, number, bathrooms, bedrooms, rooms, details, building, active FROM apartments`)

	if err != nil {
		fmt.Printf("Error getting the list of apartments. %s\n", err)
		return nil, err
	}

	var apartments = make([]*Apartment, 0)

	for rows.Next() {
		a, err := parseApartment(rows)

		if err != nil {
			fmt.Printf("Error parsing the list of apartments. %s\n", err)
			return nil, err
		}

		apartments = append(apartments, a)
	}

	return &apartments, nil
}

func (m ApartmentRepoSQL) FetchApartment(id int) (*Apartment, error) {
	rows, err := m.db.Query(`SELECT id, number, bathrooms, bedrooms, rooms, details, building, active  FROM apartments WHERE id = ?`, id)

	if err != nil {
		fmt.Printf("Error getting apartment with id %d. %s\n", id, err)
		return nil, err
	}

	for rows.Next() {
		a, err := parseApartment(rows)

		if err != nil {
			fmt.Printf("Error parsing the apartment. %s\n", err)
			return nil, err
		}

		return a, nil
	}
	return nil, nil
}

func (m ApartmentRepoSQL) PersistApartment(a *Apartment) (*Apartment, error) {
	stmt, err := m.db.Prepare(`INSERT INTO apartments (number, bathrooms, bedrooms, rooms, details, building, active) VALUES (?, ?, ?, ?, ?, ?, TRUE)`)

	if err != nil {
		fmt.Printf("Error preparing insert statement: %s\n", err)
		return nil, err
	}

	r, err := stmt.Exec(a.Number, a.Baths, a.Beds, a.Rooms, a.Details, a.Building)

	if err != nil {
		fmt.Printf("Error executing insert statement: %s\n", err)
		return nil, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		fmt.Printf("Error getting the new apartment's id. %s\n", err)
		return nil, err
	}

	a.ID = int(id)
	a.Active = true

	return a, nil

}

func (m ApartmentRepoSQL) UpdateApartment(a *Apartment) (*Apartment, error) {
	stmt, err := m.db.Prepare(`UPDATE apartments SET number = ?, bathrooms = ?, bedrooms = ?, rooms = ?, details = ?, building = ? WHERE id = ?`)

	if err != nil {
		fmt.Printf("Error preparing the update statement: %s\n", err)
		return nil, err
	}

	_, err = stmt.Exec(a.Number, a.Baths, a.Beds, a.Rooms, a.Details, a.Building, a.ID)

	if err != nil {
		fmt.Printf("Error executing update statement: %s\n", err)
		return nil, err
	}

	return a, nil
}

func (m ApartmentRepoSQL) DeleteApartment(id int) error {
	stmt, err := m.db.Prepare(`UPDATE apartments SET active = FALSE WHERE id = ?`)

	if err != nil {
		fmt.Printf("Error preparing the delete statement: %s\n", err)
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Printf("Error executing delete statement: %s\n", err)
		return err
	}

	return nil
}

func (m ApartmentRepoSQL) Close() error {
	return m.db.Close()
}

func parseApartment(rows *sql.Rows) (*Apartment, error) {
	a := new(Apartment)
	err := rows.Scan(
		&a.ID,
		&a.Number,
		&a.Baths,
		&a.Beds,
		&a.Rooms,
		&a.Details,
		&a.Building,
		&a.Active,
	)
	return a, err
}
