package main

import (
	"github.com/TBabs-codes/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	NextLocationArea *string
	PrevLocationArea *string
	word2 string
}

func main() {

	cfg := config{
		pokeapiClient:    pokeapi.NewClient(),
		NextLocationArea: nil,
		PrevLocationArea: nil,
		word2: "",
	}
	startRepl(&cfg)
}