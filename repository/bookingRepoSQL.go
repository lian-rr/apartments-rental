package repository

import (
	"database/sql"
	"fmt"
)

type BookingRepoSQL struct {
	db *sql.DB
}

func newBookingRepoSQL(db *sql.DB) BookingRepoSQL {
	return BookingRepoSQL{db: db}
}

func (m BookingRepoSQL) ListBookings() (*[]*Booking, error) {
	rows, err := m.db.Query(`SELECT id, status, startDate, endDate, details, apartment, guest, active FROM bookings`)

	if err != nil {
		fmt.Printf("Error getting the list of bookings. %s\n", err)
		return nil, err
	}

	var bookings = make([]*Booking, 0)

	for rows.Next() {
		a, err := parseBooking(rows)

		if err != nil {
			fmt.Printf("Error parsing the list of bookings. %s\n", err)
			return nil, err
		}

		bookings = append(bookings, a)
	}

	return &bookings, nil
}

func (m BookingRepoSQL) FetchBooking(id int) (*Booking, error) {
	rows, err := m.db.Query(`SELECT id, status, startDate, endDate, details, apartment, guest, active  FROM bookings WHERE id = ?`, id)

	if err != nil {
		fmt.Printf("Error getting booking with id %d. %s\n", id, err)
		return nil, err
	}

	for rows.Next() {
		a, err := parseBooking(rows)

		if err != nil {
			fmt.Printf("Error parsing the booking. %s\n", err)
			return nil, err
		}

		return a, nil
	}
	return nil, nil
}

func (m BookingRepoSQL) PersistBooking(b *Booking) (*Booking, error) {
	stmt, err := m.db.Prepare(`INSERT INTO bookings (status, startDate, endDate, details, apartment, guest, active) VALUES (?, ?, ?, ?, ?, ?, TRUE)`)

	if err != nil {
		fmt.Printf("Error preparing insert statement: %s\n", err)
		return nil, err
	}

	r, err := stmt.Exec(b.Status, b.StartDate, b.EndDate, b.Details, b.Apartment, b.Guest)

	if err != nil {
		fmt.Printf("Error executing insert statement: %s\n", err)
		return nil, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		fmt.Printf("Error getting the new booking's id. %s\n", err)
		return nil, err
	}

	b.ID = int(id)
	b.Active = true

	return b, nil

}

func (m BookingRepoSQL) UpdateBooking(b *Booking) (*Booking, error) {
	stmt, err := m.db.Prepare(`UPDATE bookings SET status = ?, startDate = ?, endDate = ?, details = ?, apartment = ?, guest = ? WHERE id = ?`)

	if err != nil {
		fmt.Printf("Error preparing the update statement: %s\n", err)
		return nil, err
	}

	_, err = stmt.Exec(b.Status, b.StartDate, b.EndDate, b.Details, b.Apartment, b.Guest, b.ID)

	if err != nil {
		fmt.Printf("Error executing update statement: %s\n", err)
		return nil, err
	}

	return b, nil
}

func (m BookingRepoSQL) DeleteBooking(id int) error {
	stmt, err := m.db.Prepare(`UPDATE bookings SET active = FALSE WHERE id = ?`)

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

func (m BookingRepoSQL) Close() error {
	return m.db.Close()
}

func parseBooking(rows *sql.Rows) (*Booking, error) {
	b := new(Booking)
	err := rows.Scan(
		&b.ID,
		&b.Status,
		&b.StartDate,
		&b.EndDate,
		&b.Details,
		&b.Apartment,
		&b.Guest,
		&b.Active,
	)
	return b, err
}
