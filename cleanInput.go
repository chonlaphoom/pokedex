package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(input string) []string {
	result := []string{}

	temp := ""
	for i, r := range input {
		if r == ' ' {
			if temp != "" {
				result = append(result, temp)
				temp = ""
			}
		} else {
			temp += string(r)
		}

		if l := i + 1; len(input) == l && temp != "" {
			result = append(result, temp)
		}
	}
	return result
}

func pokendexCMD() {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := strings.ToLower(CleanInput(text)[0])
		fmt.Printf("Your command was: %s\n", cleaned)
	}
}
