package golocation

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	id        int
	code      string
	name      string
	phonecode int
	createdAt string
	updatedAt string
	countryId int
	stateId   int
)

func AllCountries() []Country {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM countries")

	checkErr(err)

	var countries []Country

	for statement.Next() {
		err = statement.Scan(&id, &code, &name, &phonecode, &createdAt, &updatedAt)

		checkErr(err)

		country := Country{
			Id:        id,
			Code:      code,
			Name:      name,
			Phonecode: phonecode,
		}

		countries = append(countries, country)
	}

	_ = database.Close()

	return countries
}

func AllStates() []State {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM states")

	checkErr(err)

	var states []State

	for statement.Next() {
		err = statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt)

		checkErr(err)

		state := State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}

		states = append(states, state)
	}

	_ = database.Close()

	return states
}

func AllCities() []City {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM cities")

	checkErr(err)

	var cities []City

	for statement.Next() {
		err = statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt)

		checkErr(err)

		city := City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}

		cities = append(cities, city)
	}

	_ = database.Close()

	return cities
}

func GetCountry(countryId int) Country {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM countries WHERE id = ?", countryId)

	checkErr(err)

	var country Country

	defer statement.Close()

	for statement.Next() {
		err := statement.Scan(&id, &code, &name, &phonecode, &createdAt, &updatedAt)

		checkErr(err)

		country = Country{
			Id:        id,
			Code:      code,
			Name:      name,
			Phonecode: phonecode,
		}
	}

	return country
}

func GetCity(cityId int) City {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM cities WHERE id = ?", cityId)

	checkErr(err)

	var city City

	defer statement.Close()

	for statement.Next() {
		err := statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt)

		checkErr(err)

		city = City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}
	}

	return city
}

func GetState(stateId int) State {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM states WHERE id = ?", stateId)

	checkErr(err)

	var state State

	defer statement.Close()

	for statement.Next() {
		err := statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt)

		checkErr(err)

		state = State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}
	}

	return state
}

func GetCountryStates(countryId int) []State {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM states WHERE country_id = ?", countryId)

	checkErr(err)

	var states []State

	for statement.Next() {
		err = statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt)

		checkErr(err)

		state := State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}

		states = append(states, state)
	}

	_ = database.Close()

	return states
}

func GetStateCites(stateId int) []City {
	database, _ := sql.Open("sqlite3", "../location.sqlite")

	statement, err := database.Query("SELECT * FROM cities WHERE state_id = ?", stateId)

	checkErr(err)

	var cities []City

	for statement.Next() {
		err = statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt)

		checkErr(err)

		city := City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}

		cities = append(cities, city)
	}

	_ = database.Close()

	return cities
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
