package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/0x4D5352/pokedexcli/internal/config"
	"github.com/0x4D5352/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config) error
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

func ExecuteCommand(name string, cfg *config.Config) error {
	commands := getCommands()
	command, ok := commands[name]
	if !ok {
		return fmt.Errorf("error - '%s' not a valid command!", name)
	}
	err := command.callback(cfg)
	if err != nil {
		return fmt.Errorf("error - command failed with error message: %w", err)
	}
	return nil
}

func commandHelp(cfg *config.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config.Config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config.Config) error {
	err := pokeapi.GetLocation(cfg, false)
	if err != nil {
		return err
	}
	return nil
}

func commandMapBack(cfg *config.Config) error {
	if cfg.Previous == "" {
		return errors.New("already at beginning of list!")
	}
	err := pokeapi.GetLocation(cfg, true)
	if err != nil {
		return err
	}
	return nil
}
