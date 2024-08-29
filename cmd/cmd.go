package cmd

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Cmderr interface {
	Commands() map[string]cliCommand
	CommandHelp() error
	CommandMap() error
	CommandMapb() error
	CommandExit() error
}

func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the locations of the next Pokemon world list.",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the locations of the previous Pokemon world list.",
			callback:    CommandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    CommandExit,
		},
	}
}

func CommandHelp() error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	cmd := Commands()

	for _, v := range cmd {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return errors.New("Cannot call command help")
}

func CommandExit() error {

	os.Exit(1)

	return errors.New("Cannot call command exit")
}

func CommandMap() error {

	fmt.Println("Hello from Map")

	return errors.New("Cannot get a location from api")
}

func CommandMapb() error {

	fmt.Println("Hello from MapB")

	return errors.New("Cannot get a location from api")
}
