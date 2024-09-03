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
	Callback    func(*config, ...string) error
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

		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		availableCmd := Commands()
		command, ok := availableCmd[text[0]]
		if !ok {
			fmt.Println("Command does not exist.")
			continue
		}
		err := command.Callback(conf, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func Commands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"explore": {
			Name:        "explore {location _name}",
			Description: "Explore pokemons in a specific zone",
			Callback:    commandExplore,
		},
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
	}
}
