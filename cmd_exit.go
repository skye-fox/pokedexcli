package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("\nExiting...")
	os.Exit(0)
	return nil
}
