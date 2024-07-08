package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreaPokemon(areaName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		// fmt.Printf("FOUND IN CACHE | key %s\n", url)
		locationAreaResponse := RespLocationArea{}
		err := json.Unmarshal(val, &locationAreaResponse)
		if err != nil {
			return RespLocationArea{}, err
		}

		return locationAreaResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer response.Body.Close()

	byteData, err := io.ReadAll(response.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	locationAreaResponse := RespLocationArea{}
	err = json.Unmarshal(byteData, &locationAreaResponse)
	if err != nil {
		return RespLocationArea{}, err
	}

	// add an entry to the cache
	c.cache.Add(url, byteData)

	return locationAreaResponse, nil
}
