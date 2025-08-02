package main

import (
	"fmt"

	"github.com/ZacharyLozevski/pokedexcli/config"
	"github.com/ZacharyLozevski/pokedexcli/internal/pokeapi"
)

func commandExplore(config *config.Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: explore <location-area-name>")
		return nil
	}

	names, err := pokeapi.GetLocationPokemon(config, args[0])
	if err != nil {
		fmt.Printf("No such area found: %s\n", args[0])
		return nil
	}

	fmt.Printf("Exploring %s...\n", args[0])
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}