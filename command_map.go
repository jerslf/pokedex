package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.nextURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range res.Results {
		fmt.Printf("- %v\n", area.Name)
	}

	cfg.nextURL = res.Next
	cfg.previousURL = res.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousURL == nil {
		return errors.New("you're on the first page")
	}

	res, err := cfg.pokeapiClient.ListLocations(cfg.previousURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range res.Results {
		fmt.Printf("- %v\n", area.Name)
	}

	cfg.nextURL = res.Next
	cfg.previousURL = res.Previous

	return nil
}
