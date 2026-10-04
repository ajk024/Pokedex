package main

import (
	"time"

	"github.com/ajk024/Pokedex/internal/pokeapi"
)

func main() {
	initialURL := pokeapi.LocationAreaURL()
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5*time.Second, 120*time.Second), //timeout, cache interval
		nextURL:       &initialURL,
		pokedex:       pokeapi.NewPokedex(),
	}
	startRepl(cfg)
}
