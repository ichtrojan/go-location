package golocation

//Country - struct that contains the country information
type Country struct {
	Id        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Phonecode int    `json:"phonecode"`
}

//State - struct for housing the state information
type State struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CountryId int    `json:"country_id"`
}

//City - Struct for housing the city information
type City struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	StateId int    `json:"state_id"`
}
