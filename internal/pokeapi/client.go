package pokeapi

import (
	"net/http"
	"time"

	"github.com/Jidnahn/pokedexcli/internal/pokecache"
)

type PokeAPIClient struct {
	cache      *pokecache.Cache
	httpClient *http.Client
	BaseURL    string
}

func NewPokeAPIClient() *PokeAPIClient {
	return &PokeAPIClient{
		cache: pokecache.NewCache(5 * time.Minute),
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
		BaseURL: "https://pokeapi.co/api/v2",
	}
}
