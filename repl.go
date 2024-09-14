package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	running := true
	inputScanner := bufio.NewScanner(os.Stdin)
	for running {
		fmt.Print("Pokedex > ")
		inputScanner.Scan()
		input := cleanInput(inputScanner.Text())
		commandName := input[0]
		fmt.Println()
		err := executeCommand(commandName)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
