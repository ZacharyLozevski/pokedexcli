package config

import (
	"github.com/ZacharyLozevski/pokedexcli/internal/pokecache"
)

type Config struct {
  Previous string
  Next     string
  Cache    *pokecache.Cache
}