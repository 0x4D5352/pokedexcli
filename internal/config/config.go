package config

import "github.com/0x4D5352/pokedexcli/internal/pokecache"

type Config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
}
