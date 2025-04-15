package main

import (
	"fmt"
	"os"
)

func (app *App) commandExit(_ []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}
