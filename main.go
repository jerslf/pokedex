package main

import (
	"time"

	"github.com/jerslf/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: *pokeapi.NewClient(time.Hour),
	}

	startRepl(cfg)

}
