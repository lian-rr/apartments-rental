package handler

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"github.com/gorilla/mux"
	"time"
)

const dateLayout = "2006-01-02"

func StartServer() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	startHandlers(router)
	log.Fatal(http.ListenAndServe(":80", router))
}

func startHandlers(r *mux.Router) {
	fmt.Println("Starting handlers")
	r.HandleFunc("/guests", listGuests).Methods("GET")
	r.HandleFunc("/guests", addGuest).Methods("POST")
	r.HandleFunc("/guests/{id}", fetchGuest).Methods("GET")
}


func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}


type errorResp struct {
	Message string `json:"message"`
}

func sendError(m string, s int, w http.ResponseWriter) {
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(errorResp{Message: m})
}

func getParams(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func parseDate(s string) (time.Time, error) {
	t, err := time.Parse(dateLayout, s)

	if err != nil {
		fmt.Printf("Date format invalid. %s\n", err)
		return time.Time{}, err
	}

	return t, nil
}

func parseDate2String(t time.Time) string {
	return t.Format(dateLayout)
}
