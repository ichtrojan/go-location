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

func AllCountries() ([]Country, error) {

	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return nil, err
	}

	statement, err := database.Query("SELECT * FROM countries")
	if err != nil {
		return nil, err
	}

	var countries []Country

	for statement.Next() {
		err = statement.Scan(&id, &code, &name, &phonecode, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		country := Country{
			Id:        id,
			Code:      code,
			Name:      name,
			Phonecode: phonecode,
		}

		countries = append(countries, country)
	}

	_ = database.Close()

	return countries, nil
}

func AllStates() ([]State, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return nil, err
	}

	statement, err := database.Query("SELECT * FROM states")
	if err != nil {
		return nil, err
	}

	var states []State

	for statement.Next() {
		err = statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		state := State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}

		states = append(states, state)
	}

	_ = database.Close()

	return states, nil
}

func AllCities() ([]City, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return nil, err
	}

	statement, err := database.Query("SELECT * FROM cities")
	if err != nil {
		return nil, err
	}

	var cities []City

	for statement.Next() {
		err = statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		city := City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}

		cities = append(cities, city)
	}

	_ = database.Close()

	return cities, nil
}

func GetCountry(countryId int) (Country, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return Country{}, err
	}

	statement, err := database.Query("SELECT * FROM countries WHERE id = ?", countryId)
	if err != nil {
		return Country{}, err
	}

	var country Country

	defer statement.Close()

	for statement.Next() {
		err := statement.Scan(&id, &code, &name, &phonecode, &createdAt, &updatedAt)
		if err != nil {
			return Country{}, err
		}

		country = Country{
			Id:        id,
			Code:      code,
			Name:      name,
			Phonecode: phonecode,
		}
	}

	return country, nil
}

func GetCity(cityId int) (City, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return City{}, err
	}
	statement, err := database.Query("SELECT * FROM cities WHERE id = ?", cityId)
	if err != nil {
		return City{}, err
	}

	var city City

	defer statement.Close()

	for statement.Next() {
		err := statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt)
		if err != nil {
			return City{}, err
		}

		city = City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}
	}

	return city, err
}

func GetState(stateId int) (State, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return State{}, nil
	}
	statement, err := database.Query("SELECT * FROM states WHERE id = ?", stateId)
	if err != nil {
		return State{}, nil
	}

	var state State

	defer statement.Close()

	for statement.Next() {
		err := statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt)
		if err != nil {
			return State{}, nil
		}

		state = State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}
	}

	return state, nil
}

func GetCountryStates(countryId int) ([]State, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return nil, err
	}
	statement, err := database.Query("SELECT * FROM states WHERE country_id = ?", countryId)
	if err != nil {
		return nil, err
	}

	var states []State

	for statement.Next() {
		err = statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}

		state := State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}

		states = append(states, state)
	}

	_ = database.Close()

	return states, err
}

func GetStateCites(stateId int) ([]City, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return nil, err
	}
	statement, err := database.Query("SELECT * FROM cities WHERE state_id = ?", stateId)
	if err != nil {
		return nil, err
	}

	var cities []City

	for statement.Next() {
		err = statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		city := City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}

		cities = append(cities, city)
	}

	_ = database.Close()

	return cities, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
