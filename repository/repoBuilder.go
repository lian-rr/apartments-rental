package repository

import (
	"database/sql"
	"fmt"
)

const connString = "root:rootl@tcp(localhost:3308)/apartments"


func BuildGuestRepo() (GuestRepo, error) {
	return buildGuestRepoSQL()
}

func buildGuestRepoSQL() (GuestRepo, error) {

	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Printf("Error stablishing the connection with the DB: %s", err)
		return GuestRepSQL{}, err
	}

	defer db.Close()

	return newGuestRepoSQL(db), nil
}