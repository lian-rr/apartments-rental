package repository

import (
	"database/sql"
	"fmt"
)

const connString = "root:rootl@tcp(localhost:3308)/apartments?parseTime=true"


func getConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		fmt.Printf("Error stablishing the connection with the DB: %s", err)
		return nil, err
	}

	return db, err
}

func BuildGuestRepo() (GuestRepo, error) {
	return buildGuestRepoSQL()
}

func BuildBuildingRepo() (BuildingRepo, error) {
	return buildBuildingRepoSQL()
}

func BuildApartmentRepo() (ApartmentRepo, error) {
	return buildApartmentRepoSQL()
}

func buildGuestRepoSQL() (GuestRepo, error) {

	db, err := getConn()

	if err != nil {
		return GuestRepoSQL{}, err
	}

	return newGuestRepoSQL(db), nil
}

func buildBuildingRepoSQL() (BuildingRepo, error) {
	db, err := getConn()

	if err != nil {
		return BuildingRepoSQL{}, err
	}

	return newBuildingRepoSQL(db), nil
}

func buildApartmentRepoSQL() (ApartmentRepo, error) {
	db, err := getConn()

	if err != nil {
		return ApartmentRepoSQL{}, err
	}

	return newApartmentRepoSQL(db), nil
}