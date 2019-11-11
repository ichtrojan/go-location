package golocation

type Country struct {
	Id        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Phonecode int    `json:"phonecode"`
}

type State struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CountryId int    `json:"country_id"`
}

type City struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	StateId int    `json:"state_id"`
}
