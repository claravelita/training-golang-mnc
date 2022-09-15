package session6

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/claravelita/training-golang-mnc/data"
)

var PORT = ":8080"

func SessionSixMain() {
	http.HandleFunc("/employees", getEmployees)
	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		json.NewDecoder(w).Encode(data.EmployeesData)
		return
	}
	http.Error(w, "Invalid Method", http.StatusBadRequest)
}
