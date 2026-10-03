package main

import (
	"fmt"

	"github.com/ajk024/Pokedex/internal/pokeapi"
)

func commandExit(cfg *config, str string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandHelp(cfg *config, str string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range cfg.commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config, str string) error {
	return commandMapPage(cfg, cfg.nextURL, "you're on the last page")
}

func commandMapb(cfg *config, str string) error {
	return commandMapPage(cfg, cfg.previousURL, "you're on the first page")
}

func commandMapPage(cfg *config, url *string, message string) error {
	if url == nil {
		fmt.Println(message)
		return nil
	}

	res, err := cfg.pokeapiClient.Get(*url)
	return resParse(cfg, res, err, "map")
}

func commandExplore(cfg *config, area string) error {
	if area == "" { //argument was only "explore"
		fmt.Println("Input 'explore <area name>'")
		return nil
	}

	url := pokeapi.LocationAreaURL() + "/" + area

	res, err := cfg.pokeapiClient.Get(url)
	return resParse(cfg, res, err, "area")
}
