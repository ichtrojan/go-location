package golocation

import (
	"encoding/json"
)

func AllCountries() Country {
	var countries Country

	_ = json.Unmarshal(loadCountries(), &countries)

	return countries
}

func GetCountry(id int) Country {
	var country Country

	_ = json.Unmarshal(loadCountries(), &country)

	for idx, cty := range country {
		if id == country[idx].ID {
			return Country{cty}
		}
	}

	return nil
}

func AllStates() State {
	var states State

	_ = json.Unmarshal(loadStates(), &states)

	return states
}

func GetState(id int) interface{} {
	var state State

	_ = json.Unmarshal(loadStates(), &state)

	for idx, s := range state {
		if id == state[idx].ID {
			return s
		}
	}

	return nil
}

func GetStatesByCtyID(ctyid int) State {
	var state State

	_ = json.Unmarshal(loadStates(), &state)

	for i, s := range state {
		if ctyid == state[i].CountryID {
			return State{s}
		}
	}

	return nil
}

func AllCities() City {
	var cities City

	_ = json.Unmarshal(loadCities(), &cities)

	return cities
}

func GetCity(id int) interface{} {
	var city City

	_ = json.Unmarshal(loadCities(), &city)

	for idx, c := range city {
		if id == city[idx].ID {
			return c
		}
	}

	return nil
}

func GetCities(stateid int) interface{} {
	var city City

	_ = json.Unmarshal(loadCities(), &city)

	for idx, cities := range city {
		if stateid == city[idx].StateID {
			return cities
		}
	}

	return nil
}
