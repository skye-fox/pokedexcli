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
	pokedex       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type help for a list of commands.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		if cmd, ok := getCommands()[command]; ok {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("\nUnknown command: type help for a list of commands.")
			fmt.Println()
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempt to catch a pokemon!",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <area name>",
			description: "Explore an area to see what pokemon are there!",
			callback:    commandExplore,
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
