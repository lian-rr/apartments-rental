package handler

import (
	"encoding/json"
	"fmt"
	"github.com/lian-rr/apartment-rental/manager"
	"net/http"
	"strconv"
)

type Building struct {
	ID          int    `json:"id, omitempty"`
	SName       string `json:"shortName"`
	FName       string `json:"fullName"`
	Addr        string `json:"address"`
	Phone       string `json:"phone"`
	Mng         string `json:"manager"`
	Description string `json:"description"`
	Active      bool   `json:"active, omitempty"`
}

func listBuildings(w http.ResponseWriter, r *http.Request) {

	lb, err := manager.ListBuildings()

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	json.NewEncoder(w).Encode(lb)
}

func addBuilding(w http.ResponseWriter, r *http.Request) {

	var bReq Building

	err := json.NewDecoder(r.Body).Decode(&bReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	b := mCreateRB2B(&bReq)

	if err != nil {
		sendError(err.Error(), http.StatusBadRequest, w)
		return
	}

	b, err = manager.AddBuilding(b)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(mapB2Response(b))
}

func fetchBuilding(w http.ResponseWriter, r *http.Request) {

	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	b, err := manager.FetchBuilding(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if b == nil {
		sendError("Building not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapB2Response(b))
}

func updateBuilding(w http.ResponseWriter, r *http.Request) {
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	var bReq Building

	err = json.NewDecoder(r.Body).Decode(&bReq)

	if err != nil {
		fmt.Printf("Error parsing request body. %s ", err)
		sendError("Request body not valid", http.StatusBadRequest, w)
		return
	}

	b := mCreateRB2B(&bReq)
	b.ID = int(id)

	b, err = manager.UpdateBuilding(b)

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if b == nil {
		sendError("Building not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapB2Response(b))
}

func deleteBuilding(w http.ResponseWriter, r *http.Request) {
	p := getParams(r)

	id, err := strconv.ParseInt(p["id"], 10, 64)

	if err != nil {
		sendError("Invalid format for id", http.StatusBadRequest, w)
		return
	}

	b, err := manager.DeleteBuilding(int(id))

	if err != nil {
		sendError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	if b == nil {
		sendError("Building not found", http.StatusNoContent, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mapB2Response(b))
}

func mCreateRB2B(b *Building) *manager.Building {
	return &manager.Building{
		SName:       b.SName,
		FName:       b.FName,
		Addr:        b.Addr,
		Phone:       b.Phone,
		Mng:         b.Mng,
		Description: b.Description,
	}
}

func mapB2Response(b *manager.Building) *Building {
	return &Building{
		ID:          b.ID,
		SName:       b.SName,
		FName:       b.FName,
		Addr:        b.Addr,
		Phone:       b.Phone,
		Mng:         b.Mng,
		Description: b.Description,
		Active:      b.Active,
	}
}
