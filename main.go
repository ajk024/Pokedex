package main

import (
	"fmt"
	"time"

	"github.com/ajk024/Pokedex/internal/pokeapi"
)

func main() {
	initialURL := pokeapi.LocationAreaURL()

	loadPokedex, err := pokeapi.LoadPokedex(pokedexFile)
	if err != nil {
		fmt.Printf("Unable to load Pokedex.  Starting with empty Pokedex.\n%s\n", err)
		loadPokedex = pokeapi.NewPokedex()
	}

	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5*time.Second, 120*time.Second), //timeout, cache interval
		nextURL:       &initialURL,
		pokedex:       loadPokedex, //pokeapi.NewPokedex(),
	}
	startRepl(cfg)
}
