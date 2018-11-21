package handler

import (
	"encoding/json"
	"fmt"
	"github.com/lian-rr/apartments-rental/manager"
	"net/http"
	"strconv"
)

type Apartment struct {
	ID       int     `json:"id, omitempty"`
	Number   string  `json:"number"`
	Baths    float32 `json:"bathrooms"`
	Beds     int     `json:"bedrooms"`
	Rooms    int     `json:"rooms"`
	Building int     `json:"buildingId"`
	Details  string  `json:"details"`
	Active   bool    `json:"active, omitempty"`
}

func listApartments(w http.ResponseWriter, r *http.Request) {

	la, err := manager.ListApartments()

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	json.NewEncoder(w).Encode(la)
}

func addApartment(w http.ResponseWriter, r *http.Request) {

	var bReq Apartment

	err := json.NewDecoder(r.Body).Decode(&bReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	a := mCreateRA2A(&bReq)

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
		return
	}

	a, err = manager.AddApartment(a)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapA2Response(a))
}

func fetchApartment(w http.ResponseWriter, r *http.Request) {

	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	a, err := manager.FetchApartment(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if a == nil {
		sendError("Apartment not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapA2Response(a))
}

func updateApartment(w http.ResponseWriter, r *http.Request) {
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	var bReq Apartment

	err = json.NewDecoder(r.Body).Decode(&bReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	a := mCreateRA2A(&bReq)
	a.ID = int(id)

	a, err = manager.UpdateApartment(a)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if a == nil {
		sendError("Apartment not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapA2Response(a))
}

func deleteApartment(w http.ResponseWriter, r *http.Request) {
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	a, err := manager.DeleteApartment(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if a == nil {
		sendError("Apartment not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapA2Response(a))
}

func mCreateRA2A(b *Apartment) *manager.Apartment {
	return &manager.Apartment{
		Number:   b.Number,
		Baths:    b.Baths,
		Beds:     b.Beds,
		Rooms:    b.Rooms,
		Details:  b.Details,
		Building: b.Building,
	}
}

func mapA2Response(b *manager.Apartment) *Apartment {
	return &Apartment{
		ID:       b.ID,
		Number:   b.Number,
		Baths:    b.Baths,
		Beds:     b.Beds,
		Rooms:    b.Rooms,
		Details:  b.Details,
		Building: b.Building,
		Active:   b.Active,
	}
}
