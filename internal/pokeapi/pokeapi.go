package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	var data []byte
	var cacheKey string
	var ok bool

	// Determine if we navigate forwards or backwards
	if backwards {
		cacheKey = cfg.Previous
	} else {
		cacheKey = cfg.Next
	}

	if cacheKey == "" {
		// Fetch initial data if no previous next direction is set
		results, err = http.Get(baseURL + "location")
		cacheKey = baseURL + "location"
	} else {
		var cached bool
		data, cached = cfg.Cache.Get(cacheKey)
		if !cached {
			results, err = http.Get(cacheKey)
		}
	}

	if err != nil {
		return err
	}

	if results != nil {
		defer results.Body.Close()

		data, err = io.ReadAll(results.Body)
		if err != nil {
			return err
		}

		// Cache the fetched data
		cfg.Cache.Add(cacheKey, data)
	}

	if data == nil {
		log.Fatal("HOW")
	}

	contents := responseJSON{}
	err = json.Unmarshal(data, &contents)
	if err != nil {
		return err
	}

	// Update the previous and next pointers based on the response
	if contents.Previous != nil {
		cfg.Previous, ok = contents.Previous.(string)
		if !ok {
			return fmt.Errorf("unexpected type for previous field: %v", contents.Previous)
		}
	} else {
		cfg.Previous = ""
	}

	if contents.Next != nil {
		cfg.Next, ok = contents.Next.(string)
		if !ok {
			return fmt.Errorf("unexpected type for next field: %v", contents.Next)
		}
	} else {
		cfg.Next = ""
	}

	for _, result := range contents.Results {
		fmt.Println(result.Name)
	}
	return nil
}
