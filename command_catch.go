package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide the name of the Pokémon to catch")
	}

	pokemonName := args[0]

	pokemonResponse, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	maximumBaseExperience := 700                                                            // an arbitrary upper limit (blissey is 635)
	threshold := 1 - float64(pokemonResponse.BaseExperience)/float64(maximumBaseExperience) // the number to be below in order to catch
	throw := rand.Float64()                                                                 // random number, representing a throw (0 - 1)

	// if the throw is under or equal to the threshold then we caught it, otherwise it escaped
	if throw <= threshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemonResponse
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
