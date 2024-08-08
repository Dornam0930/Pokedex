package main

import (
	"fmt"
	"os"
)

type keywords struct {
	name     string
	usage    string
	callback func() error
}

func getCommands() map[string]keywords {
	return map[string]keywords{
		"help": {
			name:     "help",
			usage:    "displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			usage:    "exits the pokedex",
			callback: commandExit,
		},
		"map": {
			name:     "map",
			usage:    "displays the next 20 locations in the pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			usage:    "displays the previous 20 locations in the pokemon world",
			callback: commandMapB,
		},
	}
}

func commandHelp() error {
	fmt.Println("--------------------")
	fmt.Println("Pokedex help:")
	fmt.Println("")
	fmt.Println("Commands:")
	for _, command := range getCommands() {
		fmt.Printf("%v: %v\n", command.name, command.usage)
	}
	fmt.Println("--------------------")
	fmt.Println("")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {

	return nil
}

func commandMapB() error {

	return nil
}
