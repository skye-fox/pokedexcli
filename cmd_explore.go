package main

import "fmt"

func commandExplore(cfg *config, input []string) error {
	if len(input) < 2 || len(input) > 2 {
		fmt.Println()
		return fmt.Errorf("Invalid command, explore requires a location argument\n")
	}
	pokemonResp, err := cfg.pokeapiClient.PokemonList(input[1])
	if err != nil {
		return err
	}

	fmt.Println()
	for _, v := range pokemonResp.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
