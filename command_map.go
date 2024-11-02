package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {
	//Acquire locations from PokeAPI
	if cfg.PrevLocationArea != nil && cfg.NextLocationArea == nil {
		return fmt.Errorf("Error: End of the location areas list reached.")
	}

	
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.NextLocationArea)
	if err != nil {
		return err
	}
	//Print locations
	fmt.Println("\nLocation Areas: ")
	for _, result := range resp.Results {
		fmt.Println(" - ", result.Name)
	}
	fmt.Println()

	cfg.NextLocationArea = resp.Next
	cfg.PrevLocationArea = resp.Previous

	return nil
}

func callbackMapB(cfg *config) error {
	//Acquire locations from PokeAPI
	if cfg.PrevLocationArea == nil {
		return fmt.Errorf("Error: No previous location areas, try command - map.")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.PrevLocationArea)
	if err != nil {
		return err
	}
	//Print locations
	fmt.Println("\nLocation Areas: ")
	for _, result := range resp.Results {
		fmt.Println(" - ", result.Name)
	}
	fmt.Println()

	cfg.NextLocationArea = resp.Next
	cfg.PrevLocationArea = resp.Previous

	return nil
}
