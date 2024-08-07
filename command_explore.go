package main

import (
	"errors"
	"fmt"

	"github.com/bazmurphy/go-cli-pokedex/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("need to provide a location area name")
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	locationAreaResponse, err := cfg.pokeapiClient.ListLocationAreaPokemon(areaName)
	if err != nil {
		return err
	}

	printLocationAreaPokemon(locationAreaResponse)

	return nil
}

func printLocationAreaPokemon(locationArea pokeapi.RespLocationArea) {
	fmt.Println("Found Pokemon:")
	for _, number := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", number.Pokemon.Name)
	}
}
