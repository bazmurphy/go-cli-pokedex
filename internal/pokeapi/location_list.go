package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResponse := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResponse, nil
}
