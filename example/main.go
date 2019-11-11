package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-location"
	"net/http"
	"strconv"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		countries := golocation.AllCountries()

		_ = json.NewEncoder(w).Encode(countries)
	})

	route.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		states := golocation.AllStates()

		_ = json.NewEncoder(w).Encode(states)
	})

	route.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		cities := golocation.AllCities()

		_ = json.NewEncoder(w).Encode(cities)
	})

	route.HandleFunc("/country/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		country := golocation.GetCountry(id)

		_ = json.NewEncoder(w).Encode(country)
	})

	route.HandleFunc("/city/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		city := golocation.GetCity(id)

		_ = json.NewEncoder(w).Encode(city)
	})

	route.HandleFunc("/state/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		state := golocation.GetState(id)

		_ = json.NewEncoder(w).Encode(state)
	})

	route.HandleFunc("/states/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		states := golocation.GetCountryStates(id)

		_ = json.NewEncoder(w).Encode(states)
	})

	route.HandleFunc("/cities/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		cities := golocation.GetStateCites(id)

		_ = json.NewEncoder(w).Encode(cities)
	})

	_ = http.ListenAndServe(":9990", route)
}
