package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/ZacharyLozevski/pokedexcli/config"
	"github.com/ZacharyLozevski/pokedexcli/internal/pokeapi"
	"github.com/ZacharyLozevski/pokedexcli/models"
)

func commandCatch(config *config.Config, args []string) error {
	if config.CaughtPokemon == nil {
		config.CaughtPokemon = make(map[string]models.Pokemon)
	}

	pokemonName := args[0]

	if len(args) == 0 {
		fmt.Printf("Usage: catch <pokemon-name>")
		return nil
	}

	if _, ok := config.CaughtPokemon[pokemonName]; ok {
		fmt.Printf("This pokemon has already been caught!\n")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := pokeapi.GetPokemon(config, pokemonName)
	if err != nil {
		return err
	}

	caught := catchPokemon(pokemon.BaseExperience)
	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.CaughtPokemon[pokemonName] = *pokemon
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
}

func catchPokemon(exp int) bool {
	randNum := rand.IntN(exp)
	return randNum < 40
}