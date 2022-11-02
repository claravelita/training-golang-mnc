package session11

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"

	"github.com/claravelita/training-golang-mnc/dtos"
)

const (
	PORT                   = ":8001"
	MAX                    = 100
	MIN                    = 1
	PATH_WATER_WIND_STATUS = "assignment/session11/index.html"
)

func MainAssignmentSession11() {
	http.HandleFunc("/", StatusWinterAndWind)
	fmt.Println("application has started... on", PORT)
	http.ListenAndServe(PORT, nil)
}

func StatusWinterAndWind(w http.ResponseWriter, r *http.Request) {
	wind := helperRandomInt(MIN, MAX)
	water := helperRandomInt(MIN, MAX)
	response := dtos.Session11StatusResponse{
		Status: dtos.Session11WinterWindResponse{
			Wind:  wind,
			Water: water,
		},
	}

	template, err := template.ParseFiles(PATH_WATER_WIND_STATUS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = template.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func helperRandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
