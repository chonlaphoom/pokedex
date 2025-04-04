package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()
		cleaned := strings.ToLower(CleanInput(text)[0])
		fmt.Printf("Your command was: %s\n", cleaned)
	}
}
