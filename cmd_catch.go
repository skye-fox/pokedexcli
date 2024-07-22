package main

import (
	"fmt"
	"math/rand"

	"github.com/skye-fox/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println()
		return fmt.Errorf("Invalid command, catch requires a pokemon argument\n")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	var catchAttempt int
	if pokemon.BaseExperience > 0 {
		catchAttempt = rand.Intn(pokemon.BaseExperience)
	} else {
		fmt.Println()
		return fmt.Errorf("Invalid pokemon name\n")
	}
	catchResult := int(float64(catchAttempt) / float64(pokemon.BaseExperience) * 100)

	caught := catchPokemon(pokemon, catchResult, args...)
	fmt.Println()

	if caught {
		cfg.pokedex[pokemon.Name] = pokemon
	}

	return nil
}

func catchPokemon(pokemon pokeapi.Pokemon, catchResult int, args ...string) bool {
	fmt.Println()
	fmt.Printf("Throwing a pokeball at %s...", pokemon.Name)
	caught := false

	if pokemon.BaseExperience > 309 {
		if catchResult > 95 {
			caught = true
			fmt.Printf("\nYou caught %s!\n", args[0])
		} else {
			fmt.Printf("\n%s got away!\n", args[0])
		}
	} else if pokemon.BaseExperience > 249 {
		if catchResult > 70 {
			caught = true
			fmt.Printf("\nYou caught %s!\n", args[0])
		} else {
			fmt.Printf("\n%s got away!\n", args[0])
		}
	} else if pokemon.BaseExperience > 199 {
		if catchResult > 60 {
			caught = true
			fmt.Printf("\nYou caught %s!\n", args[0])
		} else {
			fmt.Printf("\n%s got away!\n", args[0])
		}
	} else if pokemon.BaseExperience > 149 {
		if catchResult > 50 {
			caught = true
			fmt.Printf("\nYou caught %s!\n", args[0])
		} else {
			fmt.Printf("\n%s got away!\n", args[0])
		}
	} else if pokemon.BaseExperience > 99 {
		if catchResult > 40 {
			caught = true
			fmt.Printf("\nYou caught %s!\n", args[0])
		} else {
			fmt.Printf("\n%s got away!\n", args[0])
		}
	} else if pokemon.BaseExperience > 49 {
		if catchResult > 30 {
			caught = true
			fmt.Printf("\nYou caught %s!\n", args[0])
		} else {
			fmt.Printf("\n%s got away!\n", args[0])
		}
	} else if pokemon.BaseExperience > 0 {
		if catchResult > 15 {
			caught = true
			fmt.Printf("\nYou caught %s!\n", args[0])
		} else {
			fmt.Printf("\n%s got away!\n", args[0])
		}
	}
	return caught
}
