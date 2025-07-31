package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl() {
  scanner := bufio.NewScanner(os.Stdin)

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
      err := cliCommand.callback()
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
  callback      func() error
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
    }
}