package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type PokemonInfo struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func commandCatch(input []string) error {
	if len(input) != 2 {
		return fmt.Errorf("Please provide a pokenmon name.")
	}

	name := input[1]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)

	if resp.StatusCode == 404 {
		return fmt.Errorf("Pokenmon %s not found", name)
	}

	if err != nil {
		return fmt.Errorf("failed to catch %s: %w", name, err)
	}

	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return readErr
	}

	var pokemonInfo PokemonInfo
	if err := json.Unmarshal(body, &pokemonInfo); err != nil {
		return fmt.Errorf("failed to unmarshal %w", err)
	}

	if canCatchPokemon(&pokemonInfo) {
		fmt.Printf("%s was caught!\n", name)
		state.PokemonDex[name] = pokemonInfo

		fmt.Println("inspect dex: ", state.PokemonDex)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func canCatchPokemon(pokemonInfo *PokemonInfo) bool {
	isCaught := randomBool(float64(pokemonInfo.BaseExperience / 100))
	return isCaught
}

// TODO improve weight calculation
func randomBool(weight float64) bool {
	if weight <= 1 {
		return true
	}
	return rand.Float64() < (1.0 / weight)
}
