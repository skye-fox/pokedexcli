package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
				err := cmd.callback()
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
	callback    func() error
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
	}
}
