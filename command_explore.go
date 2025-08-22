package main

import (
	"errors"
	"fmt"
	"log"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationName := args[0]

	res, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Exploring %s:", res.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
