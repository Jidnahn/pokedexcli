package pokeapi

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Config struct {
	Page          int
	Pokedex       map[string]Pokemon
	PokeAPIClient *PokeAPIClient
}
