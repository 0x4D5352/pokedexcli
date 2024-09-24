package main

import (
	"github.com/0x4D5352/pokedexcli/internal/pokeapi"
)

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  []struct {
		Name string
		Stat int
	}
	Types []string
}

func catchPokemon(pokemon pokeapi.RespPokemon) Pokemon {
	s := make([]struct {
		Name string
		Stat int
	}, len(pokemon.Stats))
	for _, stat := range pokemon.Stats {
		s = append(s, struct {
			Name string
			Stat int
		}{
			Name: stat.Stat.Name,
			Stat: stat.BaseStat,
		})
	}

	t := make([]string, len(pokemon.Types))
	for _, tp := range pokemon.Types {
		t = append(t, tp.Type.Name)
	}

	return Pokemon{
		Name:   pokemon.Name,
		Height: pokemon.Height,
		Weight: pokemon.Weight,
		Stats:  s,
		Types:  t,
	}
}
