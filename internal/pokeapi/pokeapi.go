package pokeapi

import (
	"net/http"
	"time"

	"github.com/TBabs-codes/Pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		cache: pokecache.NewCache(5*time.Minute),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
