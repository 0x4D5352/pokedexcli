package pokeapi

import (
	"github.com/0x4D5352/pokedexcli/internal/config"
	"net/http"
)

var baseURL = "https://pokeapi.co/api/v2/"

func GetLocation(cfg *config.Config, direction string) (location string, err error) {
	if direction == "" {
		results, err := http.Get(baseURL + "location")
	}
	return "", nil
}
