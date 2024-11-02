package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	// Read input line by line
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text() // Get the current line of text
		cleanText := cleanInput(text)

		//user enters in nothing return to top of loop
		if len(cleanText) == 0 {
			fmt.Print("\n")
			continue
		}

		commandName := cleanText[0]
		if len(cleanText) == 2 {
			cfg.word2 = cleanText[1]
		}
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command.")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore",
			description: "Shows Pokemon located at specified area",
			callback:    callbackExplore,
		},
		"map": {
			name:        "map",
			description: "Prints list of 20 locations in Pokemon world. \n        Subsequent calls will display next 20 locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints previous list of 20 locations in Pokemon world.",
			callback:    callbackMapB,
		},
		"help": {
			name:        "help",
			description: "Prints help menu.",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits program.",
			callback:    callbackExit,
		},
	}
}
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}
