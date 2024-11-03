package main

import "fmt"

func callbackPokedex(cfg *config) error {
	if len(cfg.pokedex.pokedex) == 0 {
		fmt.Println("No pokemon in pokedex")
	}

	for name, _ := range cfg.pokedex.pokedex {
		fmt.Printf(" - %v\n", name)
	}

	return nil
}
