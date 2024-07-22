package main

import (
	"time"

	"github.com/skye-fox/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Minute, 5*time.Second)
	cfg := &config{
		pokedex:       map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
