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
		fmt.Printf("  %s: %s\n", c.name, c.description)
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

func commandCatch(cfg *config, pokemon string) error {
	if pokemon == "" { //argument was only "catch"
		fmt.Println("Input 'catch <Pokemon name>'")
		return nil
	}

	url := pokeapi.PokemonURL() + "/" + pokemon
	res, err := cfg.pokeapiClient.Get(url)
	return resParse(cfg, res, err, "pokemon")
}

func commandInspect(cfg *config, pokemon string) error {
	if pokemon == "" {
		fmt.Println("Input 'inspect <Pokemon name>")
		return nil
	}

	val, ok := cfg.pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon, val.Height, val.Weight)
	fmt.Printf("Stats:\n")

	for _, stat := range val.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pokeType := range val.Types {
		fmt.Printf("  - %s\n", pokeType.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, str string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Printf("Your Pokedex is empty.\n")
	} else {
		fmt.Printf("Your Pokedex:\n")

		for key := range cfg.pokedex {
			fmt.Printf("  -%s\n", key)
		}
	}

	return nil
}
