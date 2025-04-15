package main

import "fmt"

func (app *App) commandInspect(input []string) error {
	if len(input) != 2 {
		return fmt.Errorf("Please provide a pokemon name.")
	}

	name := input[1]
	if pokemon, ok := app.PokemonDex[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		printPokemonStats(&pokemon)
	} else {
		fmt.Printf("Pokemon %s not found in the dex.\n", name)
	}

	return nil
}

func printPokemonStats(pokemon *PokemonInfo) {
	fmt.Printf("Stats:\n")
	for _, stats := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stats.State.Name, stats.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, stats := range pokemon.Types {
		fmt.Printf(" - %s\n", stats.Type.Name)
	}
}
