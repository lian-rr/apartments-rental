package repository

import (
	"database/sql"
	"fmt"
)

type BuildingRepoSQL struct {
	db *sql.DB
}

func newBuildingRepoSQL(db *sql.DB) BuildingRepoSQL {
	return BuildingRepoSQL{db: db}
}

func (m BuildingRepoSQL) ListBuildings() (*[]*Building, error) {
	rows, err := m.db.Query(`SELECT id, shortName, fullName, address, phone, manager, description, active FROM buildings`)

	if err != nil {
		fmt.Printf("Error getting the list of buildings. %s\n", err)
		return nil, err
	}

	var buildings = make([]*Building, 0)

	for rows.Next() {
		g, err := parseBuilding(rows)

		if err != nil {
			fmt.Printf("Error parsing the list of buildings. %s\n", err)
			return nil, err
		}

		buildings = append(buildings, g)
	}

	return &buildings, nil
}

func (m BuildingRepoSQL) FetchBuilding(id int) (*Building, error) {
	rows, err := m.db.Query(`SELECT id, shortName, fullName, address, phone, manager, description, active  FROM buildings WHERE id = ?`, id)

	if err != nil {
		fmt.Printf("Error getting building with id %d. %s\n", id, err)
		return nil, err
	}

	for rows.Next() {
		b, err := parseBuilding(rows)

		if err != nil {
			fmt.Printf("Error parsing the building. %s\n", err)
			return nil, err
		}

		return b, nil
	}
	return nil, nil
}

func (m BuildingRepoSQL) PersistBuilding(b *Building) (*Building, error) {
	stmt, err := m.db.Prepare(`INSERT INTO buildings (shortName, fullName, address, phone, manager, description, active) VALUES (?, ?, ?, ?, ?, ?, TRUE)`)

	if err != nil {
		fmt.Printf("Error preparing insert statement: %s\n", err)
		return nil, err
	}

	r, err := stmt.Exec(b.SName, b.FName, b.Addr, b.Phone, b.Mng, b.Description)

	if err != nil {
		fmt.Printf("Error executing insert statement: %s\n", err)
		return nil, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		fmt.Printf("Error getting the new building's id. %s\n", err)
		return nil, err
	}

	b.ID = int(id)
	b.Active = true

	return b, nil

}

func (m BuildingRepoSQL) UpdateBuilding(b *Building) (*Building, error) {
	stmt, err := m.db.Prepare(`UPDATE buildings SET shortName = ?, fullName = ?, address = ?, phone = ?, manager = ?, description = ? WHERE id = ?`)

	if err != nil {
		fmt.Printf("Error preparing the update statement: %s\n", err)
		return nil, err
	}

	_, err = stmt.Exec(b.SName, b.FName, b.Addr, b.Phone, b.Mng, b.Description, b.ID)

	if err != nil {
		fmt.Printf("Error executing update statement: %s\n", err)
		return nil, err
	}

	return b, nil
}

func (m BuildingRepoSQL) DeleteBuilding(id int) error {
	stmt, err := m.db.Prepare(`UPDATE buildings SET active = FALSE WHERE id = ?`)

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

func (m BuildingRepoSQL) Close() error {
	return m.db.Close()
}

func parseBuilding(rows *sql.Rows) (*Building, error) {
	b := new(Building)
	err := rows.Scan(
		&b.ID,
		&b.SName,
		&b.FName,
		&b.Addr,
		&b.Phone,
		&b.Mng,
		&b.Description,
		&b.Active,
	)
	return b, err
}
