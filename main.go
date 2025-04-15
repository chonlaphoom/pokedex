package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/chonlaphoom/pokedex/cleanInput"
	"github.com/chonlaphoom/pokedex/pokecache"
)

type App struct {
	PreviousURL string
	NextURL     string
	AppCache    *pokecache.Cache
	PokemonDex  map[string]PokemonInfo
}

func main() {
	const interval = 1 * time.Minute // reap cache every 1 minute

	app := App{
		NextURL:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PreviousURL: "",
		AppCache:    pokecache.NewCache(interval),
		PokemonDex:  make(map[string]PokemonInfo),
	}

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

			if cmd, ok := app.getCommands()[input[0]]; ok {
				err = cmd.Callback(input)
			} else {
				commandUnknown()
			}
		}

		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

type Command struct {
	Name        string
	Description string
	Callback    func(arg []string) error
}

func (app *App) getCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    app.commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Get Pokemon world map",
			Callback:    app.commandMap,
		},
		"pmap": {
			Name:        "pmap",
			Description: "Get previous Pokemon world map",
			Callback:    app.commandPMap,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the Pokemon from given area",
			Callback:    app.commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch the pokemon from given name",
			Callback:    app.commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect the pokemon from given name from pokedex",
			Callback:    app.commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Inspect pokedex",
			Callback:    app.commandPokedex,
		},
		"help": {
			Name:        "help",
			Description: "list all commands",
			Callback:    app.commandHelp,
		},
	}
}
