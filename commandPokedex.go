package main

import "fmt"

func commandPokedex(input []string) error {
	if len(state.PokemonDex) == 0 {
		fmt.Println("Your Pokemon dex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for k := range state.PokemonDex {
		fmt.Println(" - ", k)
	}

	return nil
}
