package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/chonlaphoom/pokedex/pokecache"
)

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
}

type BaseResponse[T any] struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
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

func fetchAndPrint(url string) error {
	var body []byte
	var err error

	if url == "" {
		return errors.New("URL is empty")
	}

	if val, hasCache := state.AppCache.Get(url); hasCache {
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
		state.AppCache.Add(url, _body)
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

	state.Previous = baseResponse.Previous
	state.Next = baseResponse.Next

	return nil
}

func getUrl(isPrev bool) string {
	if isPrev {
		return state.Previous
	}

	return state.Next
}
