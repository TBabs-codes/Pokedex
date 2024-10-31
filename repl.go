package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)




func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	nextPrev := config{
		next: nil,
		previous: nil,
	}
	// Read input line by line
	for  {
		fmt.Print("Pokedex command:")
		scanner.Scan()
		text := scanner.Text() // Get the current line of text
		cleanText := cleanInput(text)
		
		//user enters in nothing return to top of loop
		if len(cleanText) == 0 {
			fmt.Print("\n")
			continue
		}
		
		commandName := cleanText[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command.")
			continue
		}

		command.callback(nextPrev)

	}

}

type cliCommand struct {
	name string
	description string
	callback func(config) error
}

func getCommands() map[string]cliCommand {
	return map[string] cliCommand{
		"map": {
			name: "map",
			description: "Prints list of 20 locations in Pokemon world. \n        Subsequent calls will display next 20 locations",
			callback: callbackMap,
		},
		"help": {
			name: "help",
			description: "Prints help menu.",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits program.",
			callback: callbackExit,
		},
	}
}
func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	return strings.Fields(lowered)
}

