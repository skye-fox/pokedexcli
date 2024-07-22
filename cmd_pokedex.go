package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf(" - %v\n", k)
	}
	fmt.Println()
	return nil
}
