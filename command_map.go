package main

import (
	"errors"
	"fmt"

	"github.com/bazmurphy/go-cli-pokedex/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	printLocations(locationsResponse)

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	printLocations(locationsResponse)

	return nil
}

func printLocations(locationsResponse pokeapi.RespShallowLocations) {
	for _, location := range locationsResponse.Results {
		fmt.Println("-----", location.Name)
	}
}
