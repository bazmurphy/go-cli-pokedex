package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		// fmt.Printf("FOUND IN CACHE | key %s\n", url)
		pokemonResponse := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResponse)
		if err != nil {
			return RespPokemon{}, err
		}

		return pokemonResponse, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return RespPokemon{}, err
	}
	defer response.Body.Close()

	byteData, err := io.ReadAll(response.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResponse := RespPokemon{}
	err = json.Unmarshal(byteData, &pokemonResponse)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, byteData)

	return pokemonResponse, nil
}
