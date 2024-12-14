package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		input := sc.Text()
		clean := cleanInput(input)
		if len(clean) == 0 {
			continue
		}
		command := clean[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.TrimSpace(strings.ToLower(text)))

	return words
}
