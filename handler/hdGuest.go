package handler

import (
	"net/http"
	"encoding/json"
	"github.com/lian-rr/apartment-rental/manager"
	"fmt"
)

type Guest struct {
	id      int    `json:"id"`
	fname   string `json:"firstName"`
	lname   string `json:"lastName"`
	bdate   string `json:"birthDate"`
	gender  string `json:"gender"`
	details string `json:"details"`
	active  bool   `json:"active"`
}

func addGuest(w http.ResponseWriter, r *http.Request) {

	var gReq *Guest

	json.NewDecoder(r.Body).Decode(gReq)

	g, err := mapBody2G(gReq)

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
	}

	ng, err := manager.AddGuest(g)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapG2Body(ng))
}

func mapBody2G(g *Guest) (*manager.Guest, error) {

	bDate, err := parseDate(g.bdate)

	if err != nil {
		fmt.Printf("Birth date format not invalid.")
		return &manager.Guest{}, err
	}

	return &manager.Guest{ID: g.id, Fname: g.fname, Lname: g.lname, Bdate: bDate, Gender: g.gender, Details: g.details, Active: g.active}, nil
}

func mapG2Body(g *manager.Guest) *Guest {
	return &Guest{id: g.ID, fname: g.Fname, lname: g.Lname, bdate: parseDate2String(g.Bdate), gender: g.Gender, details: g.Details, active: g.Active}
}
