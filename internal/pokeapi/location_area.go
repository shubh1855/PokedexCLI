package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(areaName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + areaName

	if cachedData, found := c.cache.Get(url); found {
		resp := RespLocationArea{}

		err := json.Unmarshal(cachedData, &resp)
		if err != nil {
			return RespLocationArea{}, err
		}

		return resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	c.cache.Add(url, data)

	resp := RespLocationArea{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return RespLocationArea{}, err
	}

	return resp, nil
}
