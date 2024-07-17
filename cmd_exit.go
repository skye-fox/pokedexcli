package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("\nExiting...")
	os.Exit(0)
	return nil
}
