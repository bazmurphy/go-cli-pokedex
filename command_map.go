package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsResponse, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResponse.Next
	cfg.prevLocationsURL = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println("-----", location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResponse, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResponse.Next
	cfg.prevLocationsURL = locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println("-----", location.Name)
	}
	return nil
}
