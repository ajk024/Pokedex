package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		clean_str := cleanInput(line)

		if len(clean_str) > 0 {
			if cmd, ok := commands[clean_str[0]]; ok {
				if err := cmd.callback(); err != nil {
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

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading standard input: ", err)
		}
	}
	os.Exit(0)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
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
}

func cleanInput(text string) []string {
	//return []string{}
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n\n\nhelp: Displays a help message\nexit: Exit the Pokedex")
	return nil
}
