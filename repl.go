package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/ajk024/Pokedex/internal/pokeapi"
)

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
	pokedex       map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
			}
			return
		}

		line := scanner.Text()
		clean_str := cleanInput(line)

		if len(clean_str) > 0 {
			if cmd, ok := cfg.commands[clean_str[0]]; ok {
				str_arg := ""

				if len(clean_str) == 2 {
					str_arg = clean_str[1]
				}

				if err := cmd.callback(cfg, str_arg); err != nil {
					fmt.Printf("Command callback error: %v\n", err)
					break
				}
				if clean_str[0] == "exit" {
					break
				}

			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: " Displays next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area (explore <area name>)",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon! (catch <Pokemon name>)",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon (inspect <Pokemon name>)",
			callback:    commandInspect,
		},
	}
}

func resParse(cfg *config, res *http.Response, err error, method string) error {
	//fmt.Println("Entering resParse")
	if err != nil {
		return err
	}
	if res == nil {
		return fmt.Errorf("resParse res == nil")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	switch method {
	case "map":
		parseMap(cfg, body)
	case "area":
		parseArea(body)
	case "pokemon":
		parsePokemon(cfg, body)
	}
	return nil
}

func parseMap(cfg *config, body []byte) error {
	locationAreaData := pokeapi.LocationAreas{}
	if err := json.Unmarshal(body, &locationAreaData); err != nil {
		return fmt.Errorf("Error unmarshaling location area data: %s", err)
	}

	cfg.nextURL = locationAreaData.Next
	cfg.previousURL = locationAreaData.Previous

	for _, area := range locationAreaData.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func parseArea(body []byte) error {
	encounterData := pokeapi.EncounterData{}
	if err := json.Unmarshal(body, &encounterData); err != nil {
		return fmt.Errorf("Error unmarshaling encounter data: %s", err)
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", encounterData.Name)

	for _, encounter := range encounterData.PokemonEncounters {
		fmt.Printf("	- %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func parsePokemon(cfg *config, body []byte) error {
	pokemonData := pokeapi.PokemonData{}
	if err := json.Unmarshal(body, &pokemonData); err != nil {
		return fmt.Errorf("Error unmarshaling pokemon data: %s", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonData.Name)

	chance := catchChance(pokemonData.BaseExperience)
	if rand.Intn(100) < chance { //successful catch
		fmt.Printf("%s was caught!\n", pokemonData.Name)

		//Add Pokemon data into Pokedex
		pokemon := pokeapi.Pokemon{}
		if err := json.Unmarshal(body, &pokemon); err != nil {
			return fmt.Errorf("Error unmarshaling pokemon data into pokemon: %s", err)
		}
		cfg.pokedex[pokemonData.Name] = pokemon

	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}
	return nil
}

func catchChance(baseExp int) int {
	chance := 100 - baseExp/5

	//keep roll from being automatic catch or automatic failure
	if chance < 5 {
		chance = 5
	}
	if chance > 95 {
		chance = 95
	}

	return chance
}
