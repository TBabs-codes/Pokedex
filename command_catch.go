package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config) error {
	if cfg.word2 == "" {
		return fmt.Errorf("no pokemon name provided")
	}

	resp, err := cfg.pokeapiClient.PokemonReq(cfg.word2)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", cfg.word2)
	
	if rand.Intn(resp.BaseExperience) < (resp.BaseExperience/2) {
		fmt.Printf("%v escaped!\n", cfg.word2)
		return nil
	}

	fmt.Printf("%v was caught!\n", cfg.word2)

	cfg.pokedex.Add(cfg.word2, resp)

	return nil
}