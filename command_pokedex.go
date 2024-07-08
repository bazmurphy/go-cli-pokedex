package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, entry := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", entry.Name)
	}
	return nil
}
