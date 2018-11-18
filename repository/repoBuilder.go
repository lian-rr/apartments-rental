package repository

import (
	"database/sql"
	"fmt"
)

const connString = "root:rootl@tcp(localhost:3308)/apartments?parseTime=true"

func BuildGuestRepo() (GuestRepo, error) {
	return buildGuestRepoSQL()
}

func BuildBuildingRepo() (BuildingRepo, error) {
	return buildBuildingRepoSQL()
}

func buildGuestRepoSQL() (GuestRepo, error) {

	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Printf("Error stablishing the connection with the DB: %s", err)
		return GuestRepoSQL{}, err
	}

	return newGuestRepoSQL(db), nil
}

func buildBuildingRepoSQL() (BuildingRepo, error) {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Printf("Error stablishing the connection with the DB: %s", err)
		return BuildingRepoSQL{}, err
	}

	return newBuildingRepoSQL(db), nil
}
