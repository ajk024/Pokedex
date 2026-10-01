package main

import (
	"fmt"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, c := range cfg.commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	return commandMapPage(cfg, cfg.nextURL, "you're on the last page")
}

func commandMapb(cfg *config) error {
	return commandMapPage(cfg, cfg.previousURL, "you're on the first page")
}

func commandMapPage(cfg *config, url *string, message string) error {
	if url == nil {
		fmt.Println(message)
		return nil
	}

	res, err := cfg.pokeapiClient.Get(*url)
	return resParse(cfg, res, err)
}
