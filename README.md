# Pokedex CLI

A simple interactive command-line application written in Go that lets you explore the Pokémon world using the [PokeAPI](https://pokeapi.co/).  

This project was built to practice:
- Using Go’s `net/http` package for making API requests
- JSON unmarshalling into Go structs
- Building a REPL (Read–Eval–Print Loop)
- Using `sync.Mutex` for safe concurrent access to a cache
- Designing with internal packages in Go (`internal/pokeapi`, `internal/pokecache`)

---

## ✨ Features

- **help** → Show available commands  
- **map** → List the next page of location areas (20 at a time)  
- **mapb** → Go back to the previous page of location areas  
- **explore {location_area}** → Show Pokémon that can be found in a given area  
- **catch {pokemon_name}** → Attempt to catch a Pokémon (catch chance depends on base experience)  
- **inspect {pokemon_name}** → Show detailed info about a caught Pokémon (name, height, weight, stats, type)  
- **pokedex** → List all Pokémon you’ve successfully caught  
- **exit** → Quit the CLI
