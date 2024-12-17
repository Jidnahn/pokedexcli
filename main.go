package main

import (
	"github.com/Jidnahn/pokedexcli/internal/pokeapi"
)

func main() {
	config := pokeapi.Config{
		Page:          1,
		Pokedex:       map[string]pokeapi.Pokemon{},
		PokeAPIClient: pokeapi.NewPokeAPIClient(),
	}
	startRepl(&config)
}
