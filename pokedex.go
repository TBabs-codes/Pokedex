package main

import (
	"github.com/TBabs-codes/Pokedex/internal/pokeapi"
)

type Pokedex struct {
	pokedex map[string]pokeapi.PokemonResponse
}

func newPokedex() Pokedex {
	new_pokedex := Pokedex{
		pokedex: make(map[string]pokeapi.PokemonResponse),
	}

	return new_pokedex
}

func (p Pokedex) Add(pokemon string, poke_data pokeapi.PokemonResponse) {
	p.pokedex[pokemon] = poke_data
}

func (p Pokedex) Get(pokemon string) (pokeapi.PokemonResponse, bool) {
	data, ok := p.pokedex[pokemon]

	return data, ok
}