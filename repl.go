package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
				if err := cmd.callback(cfg); err != nil {
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
	}
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

	locationAreaData := pokeapi.LocationAreas{}
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
