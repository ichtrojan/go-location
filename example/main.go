package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	golocation "github.com/ichtrojan/go-location"
)

func main() {
	route := mux.NewRouter()
	golocation, err := golocation.New()
	if err != nil {
		log.Fatal(err)
	}

	route.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		countries, err := golocation.AllCountries()
		if err != nil {
			log.Fatal(err)
		}

		_ = json.NewEncoder(w).Encode(countries)
	})

	route.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		states, err := golocation.AllStates()
		if err != nil {
			log.Fatal(err)
		}

		_ = json.NewEncoder(w).Encode(states)
	})

	route.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		cities, err := golocation.AllCities()
		if err != nil {
			log.Fatal(err)
		}

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

		city, err := golocation.GetCity(id)
		if err != nil {
			log.Fatal(err)
		}

		_ = json.NewEncoder(w).Encode(city)
	})

	route.HandleFunc("/state/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])

		state, err := golocation.GetState(id)
		if err != nil {
			log.Fatal(err)
		}

		_ = json.NewEncoder(w).Encode(state)
	})

	route.HandleFunc("/states/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		states, err := golocation.GetCountryStates(id)
		if err != nil {
			log.Fatal(err)
		}

		_ = json.NewEncoder(w).Encode(states)
	})

	route.HandleFunc("/cities/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, _ := strconv.Atoi(vars["id"])

		cities, err := golocation.GetStateCites(id)
		if err != nil {
			log.Fatal(err)
		}

		_ = json.NewEncoder(w).Encode(cities)
	})

	fmt.Println("Server started")
	_ = http.ListenAndServe(":9990", route)
}
