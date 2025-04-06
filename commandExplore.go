package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Pokenmon struct {
	Name string `json:"name"`
}

type PokemonData struct {
	Pokemon Pokenmon `json:"pokemon"`
}

type LocationArea struct {
	PokemonEncounter []PokemonData `json:"pokemon_encounters"`
}

func commandExplore(input []string) error {
	if len(input) != 2 {
		return errors.New("Please provide a location name")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + input[1]
	var locationArea LocationArea

	if val, hasCache := state.AppCache.Get(url); hasCache {
		if err := json.Unmarshal(val, &locationArea); err != nil {
			fmt.Println("error from unmarshal")
			return err
		}
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return readErr
		}

		if err := json.Unmarshal(body, &locationArea); err != nil {
			fmt.Println("error from unmarshal")
			return err
		}

		state.AppCache.Add(url, body)
	}

	// print names
	for _, encounter := range locationArea.PokemonEncounter {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
