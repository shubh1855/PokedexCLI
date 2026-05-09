package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a location area")
		return nil
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)

	resp, err := cfg.pokeapiClient.ExploreLocation(areaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
