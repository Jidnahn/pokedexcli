package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Jidnahn/pokedexcli/internal/pokeapi"
)

func startRepl(config *pokeapi.Config) {
	// create scanner
	sc := bufio.NewScanner(os.Stdin)

	// available commands
	for {
		// take input
		fmt.Print("Pokedex > ")
		sc.Scan()
		input := sc.Text()
		clean := cleanInput(input)
		if len(clean) == 0 {
			continue
		}
		command := clean[0]
		// check if input is valid and execute callback
		if val, ok := getCommands(config)[command]; ok {
			if len(clean) > 1 {
				param := clean[1]
				switch command {
				case "explore", "catch", "inspect":
					if err := val.callback(config, param); err != nil {
						fmt.Println("Error:", err)
					}
				default:
					fmt.Println("Please provide a valid argument")
				}
			} else {
				err := val.callback(config)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.TrimSpace(strings.ToLower(text)))

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config, args ...string) error
}

func getCommands(config *pokeapi.Config) map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the current 20 location areas",
			callback:    commandMap,
		},
		"mapn": {
			name:        "mapn",
			description: "Display the next 20 location areas",
			callback:    commandMapn,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Display the pokemon in the selected location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon by name",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon registered in the pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the pokemon you have registered in your pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}

	return supportedCommands
}
