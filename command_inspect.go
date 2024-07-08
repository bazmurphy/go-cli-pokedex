package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("please provide a pokémon name to inspect")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokémon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, entry := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", entry.Stat.Name, entry.BaseStat)
	}
	fmt.Println("Types:")
	for _, entry := range pokemon.Types {
		fmt.Printf("- %s\n", entry.Type.Name)
	}

	return nil
}
