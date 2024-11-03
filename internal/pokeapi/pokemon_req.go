package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonReq(pokemonName string) (PokemonResponse, error) {

	fullURL := baseURL + "/pokemon/" + pokemonName

	if cacheData, ok := c.cache.Get(fullURL); ok {
		resp := PokemonResponse{}
		err := json.Unmarshal(cacheData, &resp)
		if err != nil {
			return PokemonResponse{}, err
		}

		return resp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Check spelling of pokemon name.")
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, data)

	resp := PokemonResponse{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return PokemonResponse{}, err
	}

	return resp, nil
}
