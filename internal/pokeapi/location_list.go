package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationListResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetPokeAPIMap(config *Config, action string) error {
	pageNum := config.Page
	switch action {
	case "NEXT":
		pageNum++
	case "PREV":
		pageNum--
	default:
		break
	}
	url := fmt.Sprintf("%s/location-area/?limit=20&offset=%d", config.PokeAPIClient.BaseURL, (pageNum-1)*20)
	// check if petition lives in cache
	cache := config.PokeAPIClient.cache
	if data, ok := cache.Get(url); ok {
		var results LocationListResponse
		err := json.Unmarshal(data, &results)
		if err != nil {
			return fmt.Errorf("error decoding data from cache: %w", err)
		}
		// data exists in cache and has been decoded
		for _, location := range results.Results {
			fmt.Println(location.Name)
		}
		return nil
	}
	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating map request: %w", err)
	}

	// get client from config
	client := config.PokeAPIClient.httpClient

	// do request
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making map request: %w", err)
	}
	defer res.Body.Close()

	// read response and add it to the cache
	var results LocationListResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}
	cache.Add(url, data)

	// Unmarshal the data
	if err := json.Unmarshal(data, &results); err != nil {
		return fmt.Errorf("error decoding data: %w", err)
	}

	// print the names of the locations
	for _, location := range results.Results {
		fmt.Println(location.Name)
	}
	return nil
}
