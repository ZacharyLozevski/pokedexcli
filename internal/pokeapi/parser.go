package pokeapi

import (
	"encoding/json"

	"github.com/ZacharyLozevski/pokedexcli/models"
)

type PokemonData struct {
	Next interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

type NearbyPokemonData struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func parsePokemonData(rawData []byte) (*PokemonData, error) {
	// parse the body into a PokemonData struct
	var pokemonData *PokemonData
	err := json.Unmarshal(rawData, &pokemonData)
	if err != nil {
		return nil, err
	}

	return pokemonData, nil
}

func parsePokemon(rawData []byte) (*models.Pokemon, error) {
	// parse the body into a Pokemon struct
	var pokemon *models.Pokemon
	err := json.Unmarshal(rawData, &pokemon)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func parseNearbyPokemonData(rawData []byte) (*NearbyPokemonData, error) {
	// parse the body into a NearbyPokemonData struct
	var pokemonData *NearbyPokemonData
	err := json.Unmarshal(rawData, &pokemonData)
	if err != nil {
		return nil, err
	}

	return pokemonData, nil
}

func extractPokemonNames(data *PokemonData) []string {
	var names []string
	for _, result := range data.Results {
		names = append(names, result.Name)
	}
	return names
}
