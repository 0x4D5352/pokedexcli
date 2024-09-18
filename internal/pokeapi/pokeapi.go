package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/0x4D5352/pokedexcli/internal/config"
)

var baseURL = "https://pokeapi.co/api/v2/"

type responseJSON struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(cfg *config.Config, backwards bool) error {
	var results *http.Response
	var err error
	if cfg.Next == "" {
		results, err = http.Get(baseURL + "location")
	} else if backwards {
		results, err = http.Get(cfg.Previous)
	} else {
		results, err = http.Get(cfg.Next)
	}
	if err != nil {
		return err
	}
	defer results.Body.Close()
	data, err := io.ReadAll(results.Body)
	if err != nil {
		return err
	}
	contents := responseJSON{}
	err = json.Unmarshal(data, &contents)
	if err != nil {
		return err
	}
	if contents.Previous == nil {
		cfg.Previous = ""
	} else {
		var ok bool
		cfg.Previous, ok = contents.Previous.(string)
		if !ok {
			return fmt.Errorf("unexpected type for previous field: %v", contents.Previous)
		}
	}

	if contents.Next == nil {
		cfg.Next = ""
	} else {
		var ok bool
		cfg.Next, ok = contents.Next.(string)
		if !ok {
			return fmt.Errorf("unexpected type for next field: %v", contents.Next)
		}
	}
	for _, result := range contents.Results {
		fmt.Println(result.Name)
	}
	return nil

}
