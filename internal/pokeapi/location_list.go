package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		fmt.Printf("FOUND IN CACHE | key %s\n", url)
		locationsResponse := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResponse)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResponse := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationsResponse, nil
}
