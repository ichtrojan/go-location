package golocation

import (
	"fmt"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	app, err := New()
	if err != nil {
		log.Fatal(err)
	}
	if app == nil {
		log.Fatal("Invalid App returned")
	}
}

func TestAllCountries(t *testing.T) {
	app, err := New()
	if err != nil {
		log.Fatal(err)
	}

	countries, err := app.AllCountries()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countries)
}
