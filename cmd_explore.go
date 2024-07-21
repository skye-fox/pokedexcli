package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println()
		return fmt.Errorf("Invalid command, explore requires a location argument\n")
	}
	pokemonResp, err := cfg.pokeapiClient.PokemonList(args[0])
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Exploring %s...\n", pokemonResp.Name)
	fmt.Println("Found Pokemon: ")
	for _, v := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", v.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
