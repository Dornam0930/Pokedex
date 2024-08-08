package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Errorf("Error: %v", err)
			continue
		}

		input = cleanInput(input)

		command, exists := getCommands()[input]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}
}

func cleanInput(input string) string {
	input = strings.TrimRight(input, "\n")
	input = strings.ToLower(input)
	return input
}
