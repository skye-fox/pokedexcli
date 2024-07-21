package main

import (
	"fmt"
)

func commandMap(cfg *config, input []string) error {
	locationResp, err := cfg.pokeapiClient.LocationList(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Printf("\n%v", loc.Name)
	}
	fmt.Println()
	fmt.Println()
	return nil
}

func commandMapB(cfg *config, input []string) error {
	if cfg.previous == nil {
		fmt.Println()
		return fmt.Errorf("Already on first page\n")
	}

	locationResp, err := cfg.pokeapiClient.LocationList(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Printf("\n%v", loc.Name)
	}
	fmt.Println()
	fmt.Println()
	return nil
}
