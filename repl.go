package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ZacharyLozevski/pokedexcli/config"
)

func startRepl() {
  scanner := bufio.NewScanner(os.Stdin)
  config := new(config.Config)

  for {
    // inital statement 
    fmt.Print("Pokedex > ")

    // build a scanner that reads from the terminal

    // scan for users first input
    scanner.Scan()

    // extract the first word from the input text
    input := scanner.Text()
    command := cleanInput(input)[0]

    cliCommand, ok := getCommands()[command]
    if ok {
      err := cliCommand.callback(config)
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
  callback      func(*config.Config) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand {
      "exit": {
          name: "exit",
          description: "Exit the Pokedex",
          callback: commandExit,
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