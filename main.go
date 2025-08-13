package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		input := scanner.Text()
		text := cleanInput(input)
		if len(text) == 0 {
			fmt.Println("Input was empty")
			continue
		}
		fmt.Printf("Your command was: %v\n", text[0])
	}
}
