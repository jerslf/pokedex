package main

import (
	"fmt"
	"log"

	"github.com/jerslf/pokedex/internal/pokeapi"
)

func main() {
	//startRepl()

	res, err := pokeapi.ListLocations()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
