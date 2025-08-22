package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Throwing a Pokeball at %s...:\n", pokemonName)

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("%s escaped! ", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	return nil

}
