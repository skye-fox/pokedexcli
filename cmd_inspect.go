package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Invalid command, inspect requires a pokemon argument")
	}

	fmt.Println()
	if pokemon, ok := cfg.pokedex[args[0]]; ok {
		fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" - %v\n", t.Type.Name)
		}
	} else {
		return fmt.Errorf("No entry for %v. You need to catch one first!\n", args[0])
	}
	fmt.Println()
	return nil
}
