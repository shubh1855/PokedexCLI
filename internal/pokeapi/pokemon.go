package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if cachedData, found := c.cache.Get(url); found {
		resp := Pokemon{}

		err := json.Unmarshal(cachedData, &resp)
		if err != nil {
			return Pokemon{}, err
		}

		return resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	resp := Pokemon{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return Pokemon{}, err
	}

	return resp, nil
}
