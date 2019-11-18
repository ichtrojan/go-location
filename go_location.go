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
	database  *sql.DB
)

//App - This houses the unexported database, so that it wont be tampered with outside this packaage
type App struct {
	database *sql.DB
}

//New - Create a new instance of the go-location package
func New() (*App, error) {
	database, err := sql.Open("sqlite3", "../location.sqlite")
	if err != nil {
		return nil, err
	}
	app := &App{
		database: database,
	}

	return app, nil
}

//AllCountries - Function to return the list of available countries.
func (app *App) AllCountries() ([]Country, error) {
	// database, _ := sql.Open("sqlite3", "../location.sqlite")
	//at this point, we assume that the application has been initialized successfully.
	database := app.database
	defer database.Close()

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

	return countries, nil
}

//AllStates - Function to return all the available states.
func (app *App) AllStates() ([]State, error) {
	// database, _ := sql.Open("sqlite3", "../location.sqlite")
	database := app.database
	defer database.Close()

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
	return states, err
}

//AllCities - Function to return all the available cities
func (app *App) AllCities() ([]City, error) {
	// database, _ := sql.Open("sqlite3", "../location.sqlite")
	database := app.database
	defer database.Close()

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

	return cities, err
}

//GetCountry - function to retrieve the country details
func (app *App) GetCountry(countryID int) Country {
	database := app.database
	defer database.Close()

	statement, err := database.Query("SELECT * FROM countries WHERE id = ?", countryID)

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

//GetCity - function to retrieve the city information
func (app *App) GetCity(cityID int) (*City, error) {
	database = app.database
	defer database.Close()

	statement, err := database.Query("SELECT * FROM cities WHERE id = ?", cityID)
	if err != nil {
		return nil, err
	}

	var city City

	defer statement.Close()

	for statement.Next() {
		if err := statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		city = City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}
	}

	return &city, nil
}

//GetState - Function to retrieve the states information
func (app *App) GetState(stateID int) (*State, error) {
	database := app.database
	defer database.Close()

	statement, err := database.Query("SELECT * FROM states WHERE id = ?", stateID)
	if err != nil {
		return nil, err
	}
	var state State
	defer statement.Close()

	for statement.Next() {
		if err := statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		state = State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}
	}

	return &state, nil
}

//GetCountryStates - Function to retrieve the states related to a country
func (app *App) GetCountryStates(countryID int) ([]State, error) {
	database := app.database
	statement, err := database.Query("SELECT * FROM states WHERE country_id = ?", countryID)
	if err != nil {
		return nil, err
	}
	var states []State

	for statement.Next() {
		if err = statement.Scan(&id, &name, &countryId, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		state := State{
			Id:        id,
			Name:      name,
			CountryId: countryId,
		}

		states = append(states, state)
	}

	return states, nil
}

//GetStateCites - function to retrieve the citites that are present within a state
func (app *App) GetStateCites(stateID int) ([]City, error) {
	// database, _ := sql.Open("sqlite3", "../location.sqlite")
	database = app.database

	statement, err := database.Query("SELECT * FROM cities WHERE state_id = ?", stateID)
	if err != nil {
		return nil, err
	}

	var cities []City

	for statement.Next() {
		if err = statement.Scan(&id, &name, &stateId, &createdAt, &updatedAt); err != nil {
			return nil, err
		}

		city := City{
			Id:      id,
			Name:    name,
			StateId: stateId,
		}

		cities = append(cities, city)
	}

	return cities, nil
}

func checkErr(err error) error {
	return err
}
