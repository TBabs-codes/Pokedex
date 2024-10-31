package main

import "fmt"

func callbackHelp(nextPrev config) error{
	fmt.Println("")
	fmt.Println("Welcome to Pokedex Help!!")
	fmt.Println("The available commmands are:")
	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %v: %v\n", cmd.name, cmd.description)
	}

	fmt.Println("")
	return nil
}