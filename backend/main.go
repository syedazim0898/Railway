package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type SeatResponse struct {
	Position string `json:"position"`
}

var seats = []string{
	"lower",
	"middle",
	"upper",
	"lower",
	"middle",
	"upper",
	"side lower",
	"side upper",
}

func getSeatPosition(number int) string {
	if number <= 0 {
		return "Invalid ticket number"
	}
	index := (number - 1) % len(seats)
	return seats[index]
}

func seatHandler(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number <= 0 {
		http.Error(w, "Invalid ticket number", http.StatusBadRequest)
		return
	}

	position := getSeatPosition(number)
	response := SeatResponse{Position: position}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/seat", seatHandler)
	http.ListenAndServe(":8080", nil)
}