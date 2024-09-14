package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Page forwards through locations within the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Page backwards through locations within the Pokemon world",
			callback:    commandMapBack,
		},
	}
}

func executeCommand(name string) error {
	commands := getCommands()
	command, ok := commands[name]
	if !ok {
		return fmt.Errorf("error - '%s' not a valid command!", name)
	}
	err := command.callback()
	if err != nil {
		return fmt.Errorf("error - command failed with error message: %w", err)
	}
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
