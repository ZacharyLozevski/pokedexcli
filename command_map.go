package main

import (
	"fmt"

	config "github.com/ZacharyLozevski/pokedexcli/config"
	"github.com/ZacharyLozevski/pokedexcli/internal/pokeapi"
)

func printNames(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func commandMap(config *config.Config, _ []string) error {
	url := config.Next

	if url == "" {
		fmt.Println("You cannot go forward!")
		return nil
	}

	names, err := pokeapi.GetLocationAreaData(config, url)
	if err != nil {
		return err
	}

	printNames(names)

	return nil
}

func commandMapb(config *config.Config, _ []string) error {
	url := config.Previous
	
	if url == "" {
		fmt.Println("You cannot go back!")
		return nil
	}

	names, err := pokeapi.GetLocationAreaData(config, url)	
	if err != nil {
		return err
	}

	printNames(names)

	return nil
}