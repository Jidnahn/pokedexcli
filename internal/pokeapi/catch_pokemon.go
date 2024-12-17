package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func catch(pokemon Pokemon, config *Config) {
	random := rand.Float64() * 2
	value := 1.0 - float64(pokemon.BaseExperience)/1000.0
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	caught := value > random
	if caught {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		_, ok := config.Pokedex[pokemon.Name]
		if !ok {
			config.Pokedex[pokemon.Name] = pokemon
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
}

func GetPokemonInfo(config *Config, name string) error {
	url := fmt.Sprintf("%s/pokemon/%s", config.PokeAPIClient.BaseURL, name)
	// check cache
	cache := config.PokeAPIClient.cache
	if data, ok := cache.Get(url); ok {
		var result Pokemon
		err := json.Unmarshal(data, &result)
		if err != nil {
			return fmt.Errorf("error reading data from cache: %w", err)
		}
		// data exists in cache
		catch(result, config)
		return nil
	}
	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	// get client from config
	client := config.PokeAPIClient.httpClient
	// do request
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making pokemon request: %w", err)
	}
	defer res.Body.Close()
	// read response and add to cache
	var result Pokemon
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}
	cache.Add(url, data)
	// unmarshal data
	if err := json.Unmarshal(data, &result); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}
	catch(result, config)
	return nil
}
