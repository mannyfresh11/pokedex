package main

import (
	"bufio"
	"fmt"
	"mannyfresh11/pokedex/cmd"
	"os"
)

func main() {

	fmt.Println("Please type in your pokemon: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "help":
			cmd.CommandHelp()
		case "map":
			cmd.CommandMap()
		case "mapb":
			cmd.CommandMapb()
		case "exit":
			cmd.CommandExit()
		default:
			//make api call
		}
	}

}
