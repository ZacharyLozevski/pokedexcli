package pokeapi

import (
	"github.com/ZacharyLozevski/pokedexcli/config"
)

func updateURLConfig(config *config.Config, data *PokemonData) {
	// update config.Previous
	if prev, ok := data.Previous.(string); ok {
		config.Previous = prev
	} else {
		config.Previous = ""
	}

	// update config.Next
	if next, ok := data.Next.(string); ok {
		config.Next = next
	} else {
		config.Next = ""
	}
}

func GetLocationAreaData(config *config.Config, url string) ([]string, error) {
	var rawData []byte

	// check if it is already cached
	if cache, isCached := config.Cache.Get(url); isCached {
		rawData = cache
	} else {
		// make the api request if it is not cached
		data, err := locationAreaAPIRequest(url)
		if err != nil {
			return nil, err
		}
		// update the cache with new returned data
		config.Cache.Add(url, data)

		rawData = data
	}

	// decode the returned rawData into a PokemonData struct
	pokemonData, err := parsePokemonData(rawData)
	if err != nil {
		return nil, err
	}

	// update the next and previous fields of the config
	updateURLConfig(config, pokemonData)

	// extract the names of all the areas
	names := extractPokemonNames(pokemonData)

	return names, nil
}

func GetLocationPokemon(config *config.Config, locationName string) ([]string, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName + "/"
	var rawData []byte
	
	// check if cached
	if cache, isCached := config.Cache.Get(url); isCached {
		rawData = cache
	} else {
		data, err := locationAreaAPIRequest(url)	
		if err != nil {
			return nil, err
		}

		// update the cache with the data
		config.Cache.Add(url, data)
		rawData = data
	}

	// decode the pokemon name data
	parseNearbyPokemonData(rawData)
	nearbyPokemonData, err := parseNearbyPokemonData(rawData)
	if err != nil {
		return nil, err
	}

	// extract the names of the pokemon data
	var names []string
	for _, pokemonEncounter := range nearbyPokemonData.PokemonEncounters {
		names = append(names, pokemonEncounter.Pokemon.Name)
	}

	return names, nil
}