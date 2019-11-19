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

func TestEmptyDatabaseConnection(t *testing.T) {
	var app App
	countries, err := app.AllCountries()
	if countries != nil {
		log.Fatal("Countries should be nil")
	}
	if err != nil {
		log.Fatal(err)
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
