package main

import (
	"bufio"
	"fmt"

	//"net/http"
	"os"

	"github.com/ajk024/Pokedex/internal/pokeapi"
)

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
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
