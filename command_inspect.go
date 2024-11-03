package main

import (
	"fmt"
)

func callbackInspect(cfg *config) error {
	if cfg.word2 == "" {
		return fmt.Errorf("no pokemon name provided")
	}

	if data, ok := cfg.pokedex.Get(cfg.word2); ok {
		fmt.Printf(" Name: %v\n", data.Name)
		fmt.Printf(" Height: %v\n", data.Height)
		fmt.Printf(" Weight: %v\n", data.Weight)
		fmt.Printf(" Stats: \n")
		for _, stat := range data.Stats {
			fmt.Printf("   -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf(" Types:\n")
		for _, stat := range data.Types {
			fmt.Printf("   -%v\n", stat.Type.Name)
		}

		return nil
	}

	fmt.Println("Pokemon not in pokedex. Go catch that pokemon!")
	return nil
}