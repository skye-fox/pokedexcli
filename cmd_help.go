package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("\nUsage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
