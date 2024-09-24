package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.ExploreArea(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if _, ok := cfg.pokedex[args[0]]; ok {
		if len(args) != 1 {
			return errors.New("you must provide a pokemon name")
		}
		return errors.New("already caught!")
	}
	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}
	name := pokemon.Name
	fmt.Printf("Throwing a pokeball at %s...\n", name)
	// 635 was highest base XP when writing this...
	catchChance := 1.0 - (float64(pokemon.BaseExperience) / 635.0)
	catchNumber := rand.Float32()
	if catchNumber > float32(catchChance) {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}
	fmt.Printf("%s was caught!\n", name)
	cfg.pokedex[name] = catchPokemon(pokemon)
	return nil
}
