package main

import "fmt"

func (app *App) commandHelp(input []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	fmt.Println("help: Displays a help message")
	for _, cmd := range app.getCommands() {
		text := fmt.Sprintf("%s: %s", cmd.Name, cmd.Description)
		fmt.Println(text)
	}
	return nil
}
