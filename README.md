# Pokedex CLI

A command-line Pokédex application written in Go that lets you explore the Pokémon world, discover location areas, encounter Pokémon, catch them, inspect their stats, and maintain your own personal Pokédex.

This project was built as a learning project to practice Go fundamentals, CLI development, HTTP APIs, JSON parsing, caching, and application state management.

---

## Features

### Interactive CLI (REPL)
Runs as an interactive shell where commands can be executed continuously until exit.

```bash
Pokedex >
````

### Location Navigation

Browse Pokémon location areas using paginated navigation.

Commands:

```bash
map
mapb
```

Examples:

```bash
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...
```

Go back:

```bash
Pokedex > mapb
```

---

### Explore Areas

Explore a location area and see Pokémon that can be encountered there.

Command:

```bash
explore <location-area>
```

Example:

```bash
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados
 - remoraid
```

---

### Catch Pokémon

Attempt to catch Pokémon using Pokéballs.

Catch probability is based on the Pokémon’s base experience, making stronger Pokémon harder to catch.

Command:

```bash
catch <pokemon-name>
```

Example:

```bash
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
You may now inspect it with the inspect command.
```

Failed attempt:

```bash
Pokedex > catch gyarados
Throwing a Pokeball at gyarados...
gyarados escaped!
```

---

### Inspect Caught Pokémon

Inspect detailed stats of a Pokémon you have already caught.

Command:

```bash
inspect <pokemon-name>
```

Example:

```bash
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
```

If the Pokémon has not been caught:

```bash
you have not caught that pokemon
```

---

### Personal Pokédex

View all Pokémon you have successfully caught.

Command:

```bash
pokedex
```

Example:

```bash
Pokedex > pokedex
Your Pokedex:
 - pikachu
 - squirtle
 - charmander
```

---

### Help Menu

List all available commands.

Command:

```bash
help
```

---

### Exit

Exit the application.

Command:

```bash
exit
```

---

### API Response Caching

The application includes a concurrency-safe in-memory caching layer for PokeAPI responses.

Benefits:

* Faster repeated requests
* Reduced API calls
* Instant revisits for explored locations
* Better CLI responsiveness

Implementation includes:

* sync.Mutex
* time.Ticker
* automatic cache expiration

---

## Project Structure

```text
.
├── command_catch.go
├── command_exit.go
├── command_explore.go
├── command_help.go
├── command_inspect.go
├── command_map.go
├── command_pokedex.go
├── go.mod
├── main.go
├── repl.go
├── repl_test.go
└── internal
    ├── pokeapi
    │   ├── client.go
    │   ├── location_area.go
    │   ├── location_list.go
    │   ├── pokeapi.go
    │   ├── pokemon.go
    │   └── types_locations.go
    └── pokecache
        ├── cache.go
        └── cache_test.go
```

---

## Architecture

### REPL Layer

Handles:

* user input
* command parsing
* command dispatching
* runtime application state

Files:

```text
main.go
repl.go
```

---

### Command Layer

Each CLI command is implemented independently.

Examples:

```text
command_map.go
command_catch.go
command_inspect.go
```

---

### API Layer

Responsible for:

* HTTP communication
* JSON parsing
* endpoint-specific logic
* cache integration

Files:

```text
internal/pokeapi/
```

---

### Cache Layer

Custom in-memory cache with expiration support.

Files:

```text
internal/pokecache/
```

---

## Technologies Used

* Go
* PokeAPI
* net/http
* encoding/json
* sync
* time
* math/rand

---

## Installation

Clone the repository:

```bash
git clone https://github.com/shubh1855/pokedexcli.git
cd pokedexcli
```

Install dependencies:

```bash
go mod tidy
```

Run:

```bash
go run .
```

Build executable:

```bash
go build
```

---

## Testing

Run all tests:

```bash
go test ./...
```

---

## Example Session

```bash
Pokedex > help

Welcome to the Pokedex!
Usage:

help: Displays a help message
map: Get the next page of locations
mapb: Get the previous page of locations
explore: Explore a location area
catch: Catch a pokemon
inspect: Inspect a caught pokemon
pokedex: List your caught pokemon
exit: Exit the Pokedex

Pokedex > map

Pokedex > explore pastoria-city-area

Pokedex > catch pikachu

Pokedex > inspect pikachu

Pokedex > pokedex

Pokedex > exit
```

---

## Data Source

Powered by:

[https://pokeapi.co/](https://pokeapi.co/)

---
