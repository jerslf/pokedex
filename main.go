package main

import (
	"time"

	"github.com/jerslf/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	cfg := &config{
		pokeapiClient: *pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)

}
