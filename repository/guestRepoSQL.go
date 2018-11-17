package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type GuestRepSQL struct {
	db *sql.DB
}

func newGuestRepoSQL(db *sql.DB) GuestRepSQL {
	return GuestRepSQL{db: db}
}

func (m GuestRepSQL) ListGuests() ([]*Guest, error) {

	rows, err := m.db.Query(`SELECT id, firstName, lastName, birthDay, gender, details, active FROM guest`)

	if err != nil {
		fmt.Printf("Error getting the list of guests. %s\n", err)
		return make([]*Guest, 0), err
	}

	var guests = make([]*Guest, 0)

	for rows.Next() {
		g, err := parseGuest(rows)

		if err != nil {
			fmt.Printf("Error parsing the list of guests. %s\n", err)
			return make([]*Guest, 0), err
		}

		guests = append(guests, g)
	}

	return guests, nil
}

//FindGuest the guest by Id
func (m GuestRepSQL) FindGuest(id int) (*Guest, error) {
	return &Guest{}, nil
}

//PersistGuest the guest
func (m GuestRepSQL) PersistGuest(g *Guest) (*Guest, error) {

	stmt, err := m.db.Prepare(`INSERT INTO guest (firstName, lastName, birthDay, gender, details, active) VALUES (?, ?, ?, ?, ?, TRUE)`)

	if err != nil {
		fmt.Printf("Error preparing insert statement: %s\n", err)
		return &Guest{}, err
	}

	r, err := stmt.Exec(g.Fname, g.Lname, g.Bdate, g.Gender, g.Details)

	if err != nil {
		fmt.Printf("Error executing insert statement: %s\n", err)
		return &Guest{}, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		fmt.Printf("Error getting the new guest's id. %s\n", err)
		return &Guest{}, err
	}

	g.ID = int(id)
	g.Active = true

	return g, nil

}

//UpdateGuest the guest
func (m GuestRepSQL) UpdateGuest(g *Guest) (*Guest, error) {
	return g, nil
}

//DeleteGuest the guest by Id
func (m GuestRepSQL) DeleteGuest(id int) (*Guest, error) {
	return &Guest{}, nil
}

//Close the database connection
func (m GuestRepSQL) Close() error {
	return m.db.Close()
}

func parseGuest(rows *sql.Rows) (*Guest, error) {
	g := new(Guest)
	err := rows.Scan(
		&g.ID,
		&g.Fname,
		&g.Lname,
		&g.Bdate,
		&g.Gender,
		&g.Details,
		&g.Active,
	)
	return g, err
}
