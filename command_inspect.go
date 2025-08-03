package main

import (
	"fmt"

	"github.com/ZacharyLozevski/pokedexcli/config"
)

func commandInspect(config *config.Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Usage: inspect <pokemon-name>")
		return nil
	}

	pokemonName := args[0]

	if pokemon, ok := config.CaughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats: \n")

		stats := map[string]int{}

		for _, stat := range pokemon.Stats {
			statName := stat.Stat.Name
			statData := stat.BaseStat
			stats[statName] = statData
		}

		fmt.Printf("  -hp: %d\n", stats["hp"])
		fmt.Printf("  -attack: %d\n", stats["attack"])
		fmt.Printf("  -defense: %d\n", stats["defense"])
		fmt.Printf("  -special-attack: %d\n", stats["special-attack"])
		fmt.Printf("  -special-defense: %d\n", stats["special-defense"])
		fmt.Printf("  -speed: %d\n", stats["speed"])

		fmt.Printf("Types: \n")

		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %s\n", pokemonType.Type.Name)
		}
	} else {
		fmt.Println("You haven't caught this pokemon yet!")
	}

	return nil
}