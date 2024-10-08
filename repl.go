package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/0x4D5352/pokedexcli/internal/pokeapi"
)

type config struct {
	pokedex          map[string]Pokemon
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Page backwards through locations within the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "pass the name of a location in order to explore and discover Pokemon in that area!",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "pass the name of a pokemon to try and catch it! caught pokemon are stored in your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "pass the name of a pokemon you own to view its information",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "view the contents of your pokedex",
			callback:    commandPokedex,
		},
	}
}
