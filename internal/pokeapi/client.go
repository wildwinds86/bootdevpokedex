package pokeapi

import (
	"net/http"
	"time"

	"github.com/wildwinds86/bootdevpokedex/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	Pokedex    Pokedex
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache:   pokecache.NewCache(cacheInterval),
		Pokedex: NewPokedex(),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
