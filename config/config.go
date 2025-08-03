package config

import (
	"github.com/ZacharyLozevski/pokedexcli/internal/pokecache"
	"github.com/ZacharyLozevski/pokedexcli/models"
)

type Config struct {
  Previous string
  Next     string
  Cache    *pokecache.Cache
  CaughtPokemon map[string]models.Pokemon
}