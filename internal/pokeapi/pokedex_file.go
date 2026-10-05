package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func (p Pokedex) Save(path string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshaling pokedex: %s", err)
	}
	return os.WriteFile(path, data, 0644)
}

func LoadPokedex(path string) (Pokedex, error) {
	data, err := os.ReadFile(path)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Creating new Pokedex.")
		return NewPokedex(), nil
	}

	if err != nil {
		return nil, fmt.Errorf("Error loading Pokedex: %s", err)
	}

	p := NewPokedex()
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("Error unmarshaling loaded Pokedex: %s", err)
	}
	fmt.Println("Pokedex loaded!")
	return p, nil
}
