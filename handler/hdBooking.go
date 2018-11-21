package handler

import (
	"encoding/json"
	"fmt"
	"github.com/lian-rr/apartment-rental/manager"
	"net/http"
	"strconv"
)

type Booking struct {
	ID        int    `json:"id, omitempty"`
	Status    string `json:"status"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Details   string `json:"details"`
	Apartment int    `json:"apartmentId"`
	Guest     int    `json:"guestId"`
	Active    bool   `json:"active, omitempty"`
}

func listBookings(w http.ResponseWriter, _ *http.Request) {

	lb, err := manager.ListBookings()

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	json.NewEncoder(w).Encode(lb)
}

func addBooking(w http.ResponseWriter, r *http.Request) {

	var bReq Booking

	err := json.NewDecoder(r.Body).Decode(&bReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	b, err := mCreateRBook2Book(&bReq)

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
		return
	}

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
		return
	}

	b, err = manager.AddBooking(b)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapBook2Response(b))
}

func fetchBooking(w http.ResponseWriter, r *http.Request) {

	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	b, err := manager.FetchBooking(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if b == nil {
		sendError("Booking not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapBook2Response(b))
}

func updateBooking(w http.ResponseWriter, r *http.Request) {
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	var bReq Booking

	err = json.NewDecoder(r.Body).Decode(&bReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	b, err := mCreateRBook2Book(&bReq)

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
		return
	}

	b.ID = int(id)

	b, err = manager.UpdateBooking(b)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if b == nil {
		sendError("Booking not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapBook2Response(b))
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	b, err := manager.DeleteBooking(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if b == nil {
		sendError("Booking not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapBook2Response(b))
}

func mCreateRBook2Book(b *Booking) (*manager.Booking, error) {

	bStart, err := parseDate(b.StartDate)

	if err != nil {
		fmt.Printf("Start date format not invalid.")
		return nil, err
	}

	bEnd, err := parseDate(b.EndDate)

	if err != nil {
		fmt.Printf("End date format not invalid.")
		return nil, err
	}

	return &manager.Booking{
		Status:    b.Status,
		StartDate: bStart,
		EndDate:   bEnd,
		Details:   b.Details,
		Apartment: b.Apartment,
		Guest:     b.Guest,
	}, nil
}

func mapBook2Response(b *manager.Booking) *Booking {
	return &Booking{
		ID:        b.ID,
		Status:    b.Status,
		StartDate: parseDate2String(b.StartDate),
		EndDate:   parseDate2String(b.EndDate),
		Details:   b.Details,
		Apartment: b.Apartment,
		Guest:     b.Guest,
		Active:    b.Active,
	}
}
