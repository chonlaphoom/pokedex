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

	"github.com/chonlaphoom/pokedex/pokecache"
)

var globalState = State{
	Next:     "https://pokeapi.co/api/v2/location-area",
	Previous: "",
	Cache:    &pokecache.Cache{},
}

type State struct {
	Previous string
	Next     string
	Cache    *pokecache.Cache
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

type Location struct {
	Name string `json:"name"`
	// url  string
}

type BaseResponse[T any] struct {
	Next     string `field:"next"`
	Previous string `field:"previous"`
	Results  []T    `field:"results"`
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
}

func main() {
	globalState.Cache = pokecache.NewCache(5 * 60 * 1000) // 5 mins

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var err error
		if scanner.Scan() {
			text := scanner.Text()

			if len(text) == 0 {
				continue
			}

			if er := scanner.Err(); er != nil {
				os.Exit(0)
			}

			if cmd, ok := generalRegistry[text]; ok {
				err = cmd.Callback()
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

func commandExit() error {
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

	fmt.Print("Fetching data from ", url, "\n")
	if url == "" {
		return errors.New("URL is empty")
	}

	if val, hasCache := globalState.Cache.Get(url); hasCache {
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
		globalState.Cache.Add(url, _body)
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

func commandMap() error {
	url := getUrl(false)
	err := fetchAndPrint(url)

	if err != nil {
		return err
	}

	return nil
}

func commandPMap() error {
	url := getUrl(true)
	err := fetchAndPrint(url)
	if err != nil {
		return err
	}
	return nil
}
