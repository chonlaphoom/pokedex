package main

import "fmt"

func (app *App) commandPokedex(input []string) error {
	if len(app.PokemonDex) == 0 {
		fmt.Println("Your Pokemon dex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k := range app.PokemonDex {
		fmt.Println(" - ", k)
	}

	return nil
}
