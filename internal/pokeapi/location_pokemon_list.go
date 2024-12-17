package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationPokemonListResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetPokemonFromLocation(config *Config, location string) error {
	url := fmt.Sprintf("%s/location-area/%s", config.PokeAPIClient.BaseURL, location)
	// check if data exists in cache
	cache := config.PokeAPIClient.cache
	if data, ok := cache.Get(url); ok {
		var results LocationPokemonListResponse
		err := json.Unmarshal(data, &results)
		if err != nil {
			return fmt.Errorf("error decoding data from cache: %w", err)
		}
		// data exists in cache and has been decoded
		for _, encounter := range results.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
		return nil
	}

	// data does not exist in cache, create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// get client from config
	client := config.PokeAPIClient.httpClient

	// do request
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making location pokemon request: %w", err)
	}
	defer res.Body.Close()

	// read response and add it to the cache
	var results LocationPokemonListResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}
	cache.Add(url, data)

	// unmarshal the response
	if err := json.Unmarshal(data, &results); err != nil {
		return fmt.Errorf("error decoding data: %w", err)
	}

	// print the names of the pokemon in the location
	for _, pokemon := range results.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
