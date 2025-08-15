package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range getCommands() {
		fmt.Printf("%v: %v\n", key, value.description)
	}
	fmt.Println()
	return nil
}
