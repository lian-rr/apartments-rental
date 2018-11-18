package handler

import (
	"encoding/json"
	"fmt"
	"github.com/lian-rr/apartment-rental/manager"
	"net/http"
	"strconv"
)

type Guest struct {
	ID      int    `json:"id, omitempty"`
	Fname   string `json:"firstName"`
	Lname   string `json:"lastName"`
	Bdate   string `json:"birthDate"`
	Gender  string `json:"gender"`
	Details string `json:"details"`
	Active  bool   `json:"active, omitempty"`
}

func listGuests(w http.ResponseWriter, r *http.Request){

	lg, err := manager.ListGuests()

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	json.NewEncoder(w).Encode(lg)
}

func addGuest(w http.ResponseWriter, r *http.Request) {

	var gReq Guest

	err := json.NewDecoder(r.Body).Decode(&gReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	g, err := mCreateRB2G(&gReq)

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
		return
	}

	g, err = manager.AddGuest(g)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapG2Response(g))
}

func fetchGuest(w http.ResponseWriter, r * http.Request){

	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	g, err := manager.FetchGuest(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if g == nil {
		sendError("Guest not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapG2Response(g))
}


func updateGuest(w http.ResponseWriter, r *http.Request){
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	var gReq Guest

	err = json.NewDecoder(r.Body).Decode(&gReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	g, err := mCreateRB2G(&gReq)
	g.ID = int(id)

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
		return
	}

	g, err = manager.UpdateGuest(g)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if g == nil {
		sendError("Guest not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapG2Response(g))
}

func deleteGuest(w http.ResponseWriter, r *http.Request){
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	g, err := manager.DeleteGuest(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if g == nil {
		sendError("Guest not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapG2Response(g))
}


func mCreateRB2G(g *Guest) (*manager.Guest, error) {

	bDate, err := parseDate(g.Bdate)

	if err != nil {
		fmt.Printf("Birth date format not invalid.")
		return &manager.Guest{}, err
	}

	return &manager.Guest{Fname: g.Fname, Lname: g.Lname, Bdate: bDate, Gender: g.Gender, Details: g.Details}, nil
}

func mapG2Response(g *manager.Guest) *Guest {
	return &Guest{ID: g.ID, Fname: g.Fname, Lname: g.Lname, Bdate: parseDate2String(g.Bdate), Gender: g.Gender, Details: g.Details, Active: g.Active}
}
