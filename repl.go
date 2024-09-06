package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
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
			fmt.Print(err)
		}
	}
}

func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			Name:        "catch {pokemon_name}",
			Description: "Attemt to catch a pokemon and add it to your pokedex",
			Callback:    commandCatch,
		},
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
		"inspect": {
			Name:        "inspect {pokemon_name}",
			Description: "inspect pokemon caught into pokedex",
			Callback:    commandInspect,
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
		"pokedex": {
			Name:        "pokedex",
			Description: "Display all the pokemons caught",
			Callback:    commandPokedex,
		},
	}
}
