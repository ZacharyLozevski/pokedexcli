package main

import (
	"fmt"

	"github.com/ZacharyLozevski/pokedexcli/config"
	internal "github.com/ZacharyLozevski/pokedexcli/internal/pokeapi"
)

func printNames(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func commandMap(config *config.Config) error {
	url := config.Next

	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/" 
	}

	names, err := internal.LocationAreaAPIRequest(config, url)
	if err != nil {
		return err
	}

	printNames(names)

	return nil
}

func commandMapb(config *config.Config) error {
	url := config.Previous
	
	if config.Previous == "" {
		fmt.Println("You cannot go back!")
		return nil
	}

	names, err := internal.LocationAreaAPIRequest(config, url)
	if err != nil {
		return err
	}

	printNames(names)

	return nil
}