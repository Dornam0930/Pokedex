package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type keywords struct {
	name  string
	usage string
	//callback func() error
}

func main() {
	commands := map[string]keywords{
		"help": {
			name:  "help",
			usage: "displays a help message",
			//callback: commandHelp,
		},
		"exit": {
			name:  "exit",
			usage: "exits the pokedex",
			//callback: commandExit,
		},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Errorf("Error: %v", err)
			continue
		}
		input = strings.TrimRight(input, "\n")
		if input == commands["help"].name {
			fmt.Println("--------------------")
			fmt.Println("Pokedex help:")
			fmt.Println("")
			fmt.Println("Commands:")
			fmt.Printf("%v: %v\n", commands["help"].name, commands["help"].usage)
			fmt.Printf("%v: %v\n", commands["exit"].name, commands["exit"].usage)
			fmt.Println("--------------------")
		}
		if input == commands["exit"].name {
			os.Exit(0)
		}
	}
}
