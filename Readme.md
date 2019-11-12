# Go Location ▲

![banner](https://res.cloudinary.com/ichtrojan/image/upload/v1573518254/Screenshot_2019-11-12_at_01.18.23_ioaynx.png)

## Introduction 🖖

This Package offers a simple way to get Countries, Cities and States that you may need for your Application, most especially for location dropdown.

## Getting Started 💽

### Install Package

Install this Package by running:

```bash
go get github.com/ichtrojan/go-location
```

### Import Package

```go
import "github.com/ichtrojan/go-location"
```

## Usage 🧨

This package will return the output in `structs`, you can manipulate the output however you choose. You can check the `main.go` file located in the `example` directory, the sample code in there returns the output as `JSON` to the web.

### Demo

* Clone this repo: `git clone github.com/ichtrojan/go-location`
* Change directory to example folder: `cd go-location/example`
* Run the application: `go run main.go`
* Visit application:

|Endpoint|Description|
|:------------- | :----------: |
|`/country`|return all countries|
|`/country/{id}`|return a single country by its ID|
|`/state`|return all states|
|`/state/{id}`|return a single state by its ID|
|`/states/{countryID}`|return all states in a country using the country ID|
|`/city`|return all cities|
|`/city/{id}`|return a single city by its ID|
|`/cities`|return all cities in a state using the state ID|

### Package Methods

#### Get all countries

````go
golocation.AllCountries()
````

#### Get a Country

```go
golocation.GetCountry(id)
```

> **NOTE**</br>
>`id` refers to the country ID

#### Get all states

```go
golocation.AllStates()
```

#### Get a state

```go
golocation.GetState(id)
```

> **NOTE**</br>
>`id` refers to the state ID

#### Get all states in a country

```go
golocation.GetCountryStates(id)
```

> **NOTE**</br>
>`id` refers to the country ID

#### Get all cities

```get
golocation.GetCities()
```

#### Get a city

```go
golocation.GetCity(id)
```

> **NOTE**</br>
>`id` refers to the city ID

#### Get cities in a state

```go
golocation.GetStateCites(id)
```

> **NOTE**</br>
>`id` refers to the state ID

## Contribution

Free for all, if you find an issue with the package or if a group of people somehow created a new country please send in a PR.

Danke Schön 
