package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	golocation "github.com/ichtrojan/go-location"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		// errors can be checked here to return http.StatusInternalServerError response, or any suitable response;
		countries, _ := golocation.AllCountries()

		_ = json.NewEncoder(w).Encode(countries)
	})

	route.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		// errors can be checked here to return http.StatusInternalServerError response, or any suitable response;
		states, _ := golocation.AllStates()

		_ = json.NewEncoder(w).Encode(states)
	})

	route.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		cities, _ := golocation.AllCities()

		_ = json.NewEncoder(w).Encode(cities)
	})

	route.HandleFunc("/country/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		country, _ := golocation.GetCountry(id)

		_ = json.NewEncoder(w).Encode(country)
	})

	route.HandleFunc("/city/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		city, _ := golocation.GetCity(id)

		_ = json.NewEncoder(w).Encode(city)
	})

	route.HandleFunc("/state/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		state, _ := golocation.GetState(id)

		_ = json.NewEncoder(w).Encode(state)
	})

	route.HandleFunc("/states/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		states, _ := golocation.GetCountryStates(id)

		_ = json.NewEncoder(w).Encode(states)
	})

	route.HandleFunc("/cities/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		cities, _ := golocation.GetStateCites(id)

		_ = json.NewEncoder(w).Encode(cities)
	})

	_ = http.ListenAndServe(":9990", route)
}
