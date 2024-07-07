package main

import (
	"time"

	"github.com/bazmurphy/go-cli-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
