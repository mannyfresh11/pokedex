package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *config) error
}

func startREPL(conf *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		scanner.Scan()
		text := strings.Fields(strings.ToLower(scanner.Text()))

		if len(text) == 0 {
			continue
		}

		availableCmd := Commands()
		command, ok := availableCmd[text[0]]
		if !ok {
			fmt.Println("Command does not exist.")
			continue
		}
		command.Callback(conf)
	}
}

func Commands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Display the locations of the next Pokemon world list.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the locations of the previous Pokemon world list.",
			Callback:    commandMapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
	}
}
