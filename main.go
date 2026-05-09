package main

import (
	"math/rand"
	"time"

	"github.com/shubh1855/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
		randSource:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	startRepl(cfg)
}
