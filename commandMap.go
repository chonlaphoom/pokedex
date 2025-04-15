package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
}

type BaseResponse[T any] struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []T    `json:"results"`
}

func (app *App) commandMap(_ []string) error {
	url := app.NextURL
	err := fetchAndPrint(url, app)

	if err != nil {
		return err
	}
	return nil
}

func (app *App) commandPMap(_ []string) error {
	url := app.PreviousURL
	err := fetchAndPrint(url, app)
	if err != nil {
		return err
	}
	return nil
}

func fetchAndPrint(url string, app *App) error {
	var body []byte
	var err error

	if url == "" {
		return errors.New("URL is empty")
	}

	if val, hasCache := app.AppCache.Get(url); hasCache {
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

		app.AppCache.Add(url, _body)
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

	app.PreviousURL = baseResponse.Previous
	app.NextURL = baseResponse.Next

	return nil
}
