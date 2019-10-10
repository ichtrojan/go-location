package golocation

import (
	"fmt"
	"io/ioutil"
	"os"
)

func loadCountries() []byte {
	file, err := os.Open("/home/Youngestev/go/bin/src/github.com/Youngestdev/go-location/db/countries.json")

	if err != nil {
		fmt.Println(err)
	}

	countries, _ := ioutil.ReadAll(file)
	return countries
}

func loadStates() []byte {
	file, err := os.Open("/home/Youngestev/go/bin/src/github.com/Youngestdev/go-location/db/states.json")

	if err != nil {
		fmt.Println(err)
	}

	states, _ := ioutil.ReadAll(file)
	return states
}

func loadCities() []byte {
	file, err := os.Open("/home/Youngestev/go/bin/src/github.com/Youngestdev/go-location/db/cities.json")

	if err != nil {
		fmt.Println(err)
	}

	cities, _ := ioutil.ReadAll(file)
	return cities
}