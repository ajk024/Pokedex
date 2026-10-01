package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ajk024/Pokedex/internal/pokeapi"
	//"log"
	//"time"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	initialURL := "https://pokeapi.co/api/v2/location-area/"
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		nextURL:       &initialURL,
		previousURL:   nil,
	}
	startRepl(cfg)
	os.Exit(0)
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

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
	if cfg.nextURL == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	res, err := cfg.pokeapiClient.Get(*cfg.nextURL)
	return resParse(cfg, res, err)
}

func commandMapb(cfg *config) error {
	if cfg.previousURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := cfg.pokeapiClient.Get(*cfg.previousURL)
	return resParse(cfg, res, err)
}

func resParse(cfg *config, res *http.Response, err error) error {
	if err != nil {
		return fmt.Errorf("Error retrieving Pokemap location areas: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading Pokemap location areas: %s", err)
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}

	locationAreaData := locationAreas{}
	if err := json.Unmarshal(body, &locationAreaData); err != nil {
		return fmt.Errorf("Error unmarshalling data: %s", err)
	}

	cfg.nextURL = locationAreaData.Next
	cfg.previousURL = locationAreaData.Previous

	for _, c := range locationAreaData.Results {
		fmt.Println(c.Name)
	}

	return nil
}

type locationAreas struct {
	//Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		//URL  string `json:"url"`
	} `json:"results"`
}
