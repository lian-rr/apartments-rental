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

//List of guests
func (m GuestRepSQL) ListGuests() (*[]*Guest, error) {

	rows, err := m.db.Query(`SELECT id, firstName, lastName, birthDay, gender, details, active FROM guest`)

	if err != nil {
		fmt.Printf("Error getting the list of guests. %s\n", err)
		return nil, err
	}

	var guests = make([]*Guest, 0)

	for rows.Next() {
		g, err := parseGuest(rows)

		if err != nil {
			fmt.Printf("Error parsing the list of guests. %s\n", err)
			return nil, err
		}

		guests = append(guests, g)
	}

	return &guests, nil
}

//Fetch the guest by Id
func (m GuestRepSQL) FetchGuest(id int) (*Guest, error) {

	rows, err := m.db.Query(`SELECT id, firstName, lastName, birthDay, gender, details, active FROM guest WHERE id = ?`, id)

	if err != nil {
		fmt.Printf("Error getting guest with id %d. %s\n", id, err)
		return nil, err
	}

	for rows.Next() {
		g, err := parseGuest(rows)

		if err != nil {
			fmt.Printf("Error parsing the list of guests. %s\n", err)
			return nil, err
		}

		return g, nil
	}
	return nil, nil

}

//Persist the guest
func (m GuestRepSQL) PersistGuest(g *Guest) (*Guest, error) {

	stmt, err := m.db.Prepare(`INSERT INTO guest (firstName, lastName, birthDay, gender, details, active) VALUES (?, ?, ?, ?, ?, TRUE)`)

	if err != nil {
		fmt.Printf("Error preparing insert statement: %s\n", err)
		return nil, err
	}

	r, err := stmt.Exec(g.Fname, g.Lname, g.Bdate, g.Gender, g.Details)

	if err != nil {
		fmt.Printf("Error executing insert statement: %s\n", err)
		return nil, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		fmt.Printf("Error getting the new guest's id. %s\n", err)
		return nil, err
	}

	g.ID = int(id)
	g.Active = true

	return g, nil

}

//Update the guest
func (m GuestRepSQL) UpdateGuest(g *Guest) (*Guest, error) {

	stmt, err := m.db.Prepare(`UPDATE guest SET firstName = ?, lastName = ?, birthDay = ?, gender = ?, details = ? WHERE id = ?`)

	if err != nil {
		fmt.Printf("Error preparing the update statement: %s\n", err)
		return nil, err
	}

	_, err = stmt.Exec(g.Fname, g.Lname, g.Bdate, g.Gender, g.Details, g.ID)

	if err != nil {
		fmt.Printf("Error executing update statement: %s\n", err)
		return nil, err
	}

	return g, nil
}

//Delete the guest by Id
func (m GuestRepSQL) DeleteGuest(id int) error {
	stmt, err := m.db.Prepare(`UPDATE guest SET active = FALSE WHERE id = ?`)

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
