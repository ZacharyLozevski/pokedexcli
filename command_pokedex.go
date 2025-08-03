package main

import (
	"fmt"

	"github.com/ZacharyLozevski/pokedexcli/config"
)

func commandPokedex(config *config.Config, args []string) error {
	fmt.Println("Your Pokedex: ")
	for pokemon := range config.CaughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}