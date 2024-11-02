package main

import "fmt"

func callbackExplore(cfg *config) error {
	resp, err := cfg.pokeapiClient.ExploreLocationArea(cfg.word2)
	if err != nil {
		return err
	}

	fmt.Println("Exploring ", cfg.word2, "...")
	fmt.Println("Found Pokemon:")
	for _, poke := range resp.PokemonEncounters {
		fmt.Println(" - ", poke.Pokemon.Name)
	}

	return nil
}