package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, input []string) error {
	fmt.Println("\nExiting...")
	os.Exit(0)
	return nil
}
