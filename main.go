package main

import (
	"time"

	"github.com/ajk024/Pokedex/internal/pokeapi"
)

func main() {
	initialURL := pokeapi.LocationAreaURL()
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5*time.Second, 30*time.Second), //timeout, cache interval
		nextURL:       &initialURL,
	}
	startRepl(cfg)
}
