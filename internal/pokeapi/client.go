package pokeapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ajk024/Pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokecache.NewCache(interval),
	}
}

func (cli Client) Get(url string) (*http.Response, error) {
	//check if url entry is in pokeCache

	if cli.pokeCache != nil { //used to avoid panic should pokeCache not be initialized
		val, ok := cli.pokeCache.Get(url)

		if ok == false { //not in pokeCache
			fmt.Println("Contacting Poke API")

			res, errGet := cli.httpClient.Get(url)
			if errGet != nil {
				return &http.Response{}, fmt.Errorf("Error retrieving Pokemap location areas: %s", errGet)
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return &http.Response{}, fmt.Errorf("Error reading Pokemap location areas: %s", err)
			}
			if res.StatusCode > 299 {
				return &http.Response{}, fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
			}

			cli.pokeCache.Add(url, body) //add entry to pokeCache
			res.Body = io.NopCloser(bytes.NewReader(body))
			return res, nil
		}

		//val is []bytes in pokeCache.  Return http.Response
		fmt.Println("		Entry in pokeCache!!")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(val)),
		}, nil

	}

	return &http.Response{}, fmt.Errorf("ERROR")
}
