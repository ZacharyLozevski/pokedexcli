package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ZacharyLozevski/pokedexcli/config"
	"github.com/ZacharyLozevski/pokedexcli/internal/pokecache"
)

func startRepl() {
  scanner := bufio.NewScanner(os.Stdin)
  config := &config.Config{
    Next: "https://pokeapi.co/api/v2/location-area/",
    Cache: pokecache.NewCache(5 * time.Minute),
  }

  for {
    // inital statement 
    fmt.Print("Pokedex > ")

    // build a scanner that reads from the terminal

    // scan for users first input
    scanner.Scan()

    // extract the first word from the input text
    input := scanner.Text()
    commands := cleanInput(input)
    command := commands[0]

    // extract the args 
    var args []string
    if len(commands) > 1 {
      args = cleanInput(input)[1:]
    }

    cliCommand, ok := getCommands()[command]
    if ok {
      err := cliCommand.callback(config, args)
	  if err != nil {
		fmt.Println(err)
	  }
	  continue
    } else {
      fmt.Println("Unknown command")
	  continue
    }
  }
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
  name          string
  description   string
  callback      func(*config.Config, []string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand {
      "exit": {
          name: "exit",
          description: "Exit the Pokedex",
          callback: commandExit,
      },
      "explore": {
          name: "explore",
          description: "Shows nearby Pokemon in the area",
          callback: commandExplore,
      },
      "help": {
        name: "help",
        description: "Display Help information",
        callback: commandHelp,
      },
      "map": {
        name: "map",
        description: "Displays 20 location areas around the world",
        callback: commandMap,
      },
      "mapb" : {
        name: "mapb",
        description: "Displays the previous 20 location areas around the world",
        callback: commandMapb,
      },
    }

}