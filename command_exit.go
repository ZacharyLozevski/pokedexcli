package main

import (
	"fmt"
	"os"

	"github.com/ZacharyLozevski/pokedexcli/config"
)

func commandExit(*config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}