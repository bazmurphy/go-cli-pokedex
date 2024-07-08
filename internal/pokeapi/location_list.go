package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check the cache for an entry
	if val, ok := c.cache.Get(url); ok {
		// fmt.Printf("FOUND IN CACHE | key %s\n", url)
		locationsResponse := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResponse)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer response.Body.Close()

	byteData, err := io.ReadAll(response.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResponse := RespShallowLocations{}
	err = json.Unmarshal(byteData, &locationsResponse)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// add an entry to the cache
	c.cache.Add(url, byteData)

	return locationsResponse, nil
}
