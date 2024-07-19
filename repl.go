package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skye-fox/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}

func startRepl(cfg *config) {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type help for a list of commands.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedexcli > ")

		if scanner.Scan() {
			input := scanner.Text()
			input = strings.ToLower(input)

			if len(input) == 0 {
				continue
			}

			if cmd, ok := getCommands()[input]; ok {
				err := cmd.callback(cfg)
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("\nUnknown command: type help for a list of commands.")
				fmt.Println()
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Page forward through locations in the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Page backward through locations in the pokemon world",
			callback:    commandMapB,
		},
	}
}
