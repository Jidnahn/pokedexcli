package main

import (
	"fmt"
	"os"

	"github.com/Jidnahn/pokedexcli/internal/pokeapi"
)

// callbacks
func commandExit(config *pokeapi.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	commands := getCommands(config)
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *pokeapi.Config, args ...string) error {
	if err := pokeapi.GetPokeAPIMap(config, "CURR"); err != nil {
		return err
	}
	return nil
}

func commandMapb(config *pokeapi.Config, args ...string) error {
	if config.Page == 1 {
		return fmt.Errorf("current page is the first page")
	}
	if err := pokeapi.GetPokeAPIMap(config, "PREV"); err != nil {
		return err
	}
	config.Page--
	return nil
}

func commandMapn(config *pokeapi.Config, args ...string) error {
	if err := pokeapi.GetPokeAPIMap(config, "NEXT"); err != nil {
		return err
	}
	config.Page++
	return nil
}

func commandExplore(config *pokeapi.Config, args ...string) error {
	location := args[0]
	if err := pokeapi.GetPokemonFromLocation(config, location); err != nil {
		return err
	}
	return nil
}

func commandCatch(config *pokeapi.Config, args ...string) error {
	pokemon := args[0]
	if err := pokeapi.GetPokemonInfo(config, pokemon); err != nil {
		return err
	}
	return nil
}

func commandInspect(config *pokeapi.Config, args ...string) error {
	name := args[0]
	pokemon, ok := config.Pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(config *pokeapi.Config, args ...string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("Your pokedex is currently empty! Try catching some pokemon with the catch command.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key, _ := range config.Pokedex {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}
