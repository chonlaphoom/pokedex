package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/chonlaphoom/pokedex/cleanInput"
	"github.com/chonlaphoom/pokedex/pokecache"
)

type State struct {
	Previous   string
	Next       string
	AppCache   *pokecache.Cache
	PokemonDex map[string]PokemonInfo
}

var state = State{
	Next:       "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	Previous:   "",
	AppCache:   &pokecache.Cache{},
	PokemonDex: make(map[string]PokemonInfo),
}

var generalRegistry = map[string]cliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	},
	"map": {
		Name:        "map",
		Description: "Get the map of the Pokemon world",
		Callback:    commandMap,
	},
	"pmap": {
		Name:        "pmap",
		Description: "Get the map of previois Pokemon world",
		Callback:    commandPMap,
	},
	"explore": {
		Name:        "explore",
		Description: "Explore the Pokemon from given area",
		Callback:    commandExplore,
	},
	"catch": {
		Name:        "catch",
		Description: "Catch the pokemon from given name",
		Callback:    commandCatch,
	},
}

func main() {
	const interval = 1 * time.Minute
	state.AppCache = pokecache.NewCache(interval) // 5 mins

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var err error
		if scanner.Scan() {
			text := scanner.Text()
			input := cleaninput.CleanInput(text)

			if len(text) == 0 {
				continue
			}

			if er := scanner.Err(); er != nil {
				os.Exit(0)
			}

			if cmd, ok := generalRegistry[input[0]]; ok {
				err = cmd.Callback(input)
				if err != nil {
					fmt.Println("Error:", err)
				}

			} else if text == "help" {
				err = commandHelp()
			} else {
				commandUnknown()
			}
		}

		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
