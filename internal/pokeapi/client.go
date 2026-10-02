package pokeapi

import (
	"net/http"
	"time"

	"github.com/ajk024/Pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(5 * time.Second),
	}
}

func (c Client) Get(url string) (*http.Response, error) {

	/*
		if c.pokeCache != nil {
			if val, ok := c.pokeCache.G
		}
	*/

	return c.httpClient.Get(url)
}
