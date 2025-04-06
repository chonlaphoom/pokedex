package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chonlaphoom/pokedex/cleanInput"
	"github.com/chonlaphoom/pokedex/pokecache"
)

var globalState = State{
	Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	Previous: "",
	AppCache: &pokecache.Cache{},
}

type State struct {
	Previous string
	Next     string
	AppCache *pokecache.Cache
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(input []string) error
}

type Location struct {
	Name string `json:"name"`
	// url  string
}

type BaseResponse[T any] struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
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
}

func main() {
	const interval = 5 * time.Minute
	globalState.AppCache = pokecache.NewCache(interval) // 5 mins

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
					log.Fatal(err)
				}

			} else if text == "help" {
				err = commandHelp()
			}
		} else {
			commandUnknown()
		}
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func commandUnknown() {
	fmt.Println("Unknown command")
}

func commandExit(_ []string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	fmt.Println("help: Displays a help message")
	for _, cmd := range generalRegistry {
		text := fmt.Sprintf("%s: %s", cmd.Name, cmd.Description)
		fmt.Println(text)
	}
	return nil
}

func getUrl(isPrev bool) string {
	if isPrev {
		return globalState.Previous
	}

	return globalState.Next
}

func fetchAndPrint(url string) error {
	var body []byte
	var err error

	if url == "" {
		return errors.New("URL is empty")
	}

	if val, hasCache := globalState.AppCache.Get(url); hasCache {
		body = val
	} else {
		res, errorFromGet := http.Get(url)
		if errorFromGet != nil {
			return errorFromGet
		}
		defer res.Body.Close()

		_body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return readErr
		}
		body = _body
		globalState.AppCache.Add(url, _body)
	}

	var baseResponse BaseResponse[Location] = BaseResponse[Location]{}
	err = json.Unmarshal(body, &baseResponse)
	if err != nil {
		return err
	}

	locations := baseResponse.Results

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	globalState.Previous = baseResponse.Previous
	globalState.Next = baseResponse.Next

	return nil
}

func commandMap(_ []string) error {
	url := getUrl(false)
	err := fetchAndPrint(url)

	if err != nil {
		return err
	}
	return nil
}

func commandPMap(_ []string) error {
	url := getUrl(true)
	err := fetchAndPrint(url)
	if err != nil {
		return err
	}
	return nil
}

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
	fmt.Println(url)

	// TODO consume cache
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	var locationArea LocationArea
	if err := json.Unmarshal(body, &locationArea); err != nil {
		fmt.Println("error from unmarshal")
		return err
	}

	for _, encounter := range locationArea.PokemonEncounter {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
