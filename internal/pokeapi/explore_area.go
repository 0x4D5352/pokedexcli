package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(pageURL string) (RespSpecificLocation, error) {
	url := baseURL + "/location-area/" + pageURL

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespSpecificLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespSpecificLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespSpecificLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespSpecificLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespSpecificLocation{}, err
	}

	locationsResp := RespSpecificLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespSpecificLocation{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
