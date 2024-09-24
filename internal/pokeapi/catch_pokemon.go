package pokeapi

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
)

func (c *Client) CatchPokemon(pageURL string) (string, error) {
	pokemon, err := getPokemon(c, pageURL)
	if err != nil {
		return "", err
	}
	chanceToCatch := 1.0 - pokemon.BaseExperience

	return pokemon.Name, nil
}
func getPokemon(c *Client, pageURL string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pageURL

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
