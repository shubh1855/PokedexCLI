package main

import (
	"fmt"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name")
		return nil
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	catchChance := 500 - pokemon.BaseExperience
	if catchChance < 40 {
		catchChance = 40
	}

	if cfg.randSource.Intn(500) < catchChance {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
