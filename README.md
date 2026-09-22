# Pokedex CLI

A command-line REPL interface written in Go that interacts with the PokeAPI to explore locations, catch Pokemon and manage your collection. It include a custom thread-safe in-memory caching system to minimize unnecessary network calls.

## Features

 - *Interactive REPL*: Command-line loop to explore region and encounter Wild Pokemon.

 - *PokeAPI Intergration*: Fetchs real-time location and pokemmon data form the PokeAPI.

 - *Thread-Safe Caching*: Custom cache using Go mutexes and ticker-based cleanup to store HTTP responses.

 - *Pokedex Storage*: Catch Pokemon with catch rates scaled to base experience, and inspect their stats and types.

## Available Commands

 - help: Displays a list of available commands.
 - map: Displays the next 20 location areas.
 - mab: Displays the previous 20 location areas.
 - explore <area_name>: Lists all Pokemon found in a specified area.
 - catch <pokemon_name>: Attempts to catch a Pokemon.
 - inspect <pokemon_name> Displays details (name, height, weight, stats, types) of a caught Pokemon.. pokedex: List all caught Pokemon.
 - exit: Exits the program.

## Getting Started

### Prerequisiutes

 - Go (1.20 or newer version)

### Installation

1. clone the repository:
```bash
git clone https://github.com/wheel-s/pokedexcli
```
2. Download any module dependencies:
```bash
go mod download
```
### Running the Application

Run directly 
```bash
go run .
```
Or build directly
```bash
go build -o pokedex ./pokedex
```

### Running Tests

Run the tests suite ( such as cache tests):
```bash
go test ./...
```
