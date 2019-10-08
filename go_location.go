package go_location

import (
	"fmt"
	"io/ioutil"
	"os"
)

func AllCountries() []byte {
	countries, err := os.Open("./db/countries.json")

	if err != nil {
		fmt.Println(err)
	}

	countriesJson, _ := ioutil.ReadAll(countries)
	return countriesJson
}
