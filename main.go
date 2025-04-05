package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var generalRegistry = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			text := scanner.Text()

			if len(text) == 0 {
				continue
			}

			if er := scanner.Err(); er != nil {
				os.Exit(0)
			}

			if cmd, ok := generalRegistry[text]; ok {
				err := cmd.callback()
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else if text == "help" {
				fmt.Println("Welcome to the Pokedex!")
				fmt.Println("Usage:")
				fmt.Println("")

				fmt.Println("help: Displays a help message")
				for _, cmd := range generalRegistry {
					text := fmt.Sprintf("%s: %s", cmd.name, cmd.description)
					fmt.Println(text)
				}
			}
		} else {
			unknownCommand()
			continue
		}
	}
}

func unknownCommand() {
	fmt.Println("Unknown command")
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}
