package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var cacheInterval = time.Hour

func commandHelp(cong *config, args ...string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	cmd := Commands()

	for _, c := range cmd {
		fmt.Printf("- %s: %s\n", c.Name, c.Description)
	}

	fmt.Println("")

	return nil
}

func commandExit(conf *config, args ...string) error {

	os.Exit(0)

	return nil
}

func commandMap(conf *config, args ...string) error {

	resp, err := conf.client.GetLocation(conf.NextURL, cacheInterval)
	if err != nil {
		fmt.Errorf("Failed to get next location. error:%v", err)
	}

	for _, loc := range resp.Results {
		fmt.Printf("---%s\n", loc.Name)
	}

	conf.NextURL = resp.Next
	conf.PrevURL = resp.Previous

	return nil
}

func commandMapb(conf *config, args ...string) error {

	resp, err := conf.client.GetLocation(conf.PrevURL, cacheInterval)
	if err != nil {
		fmt.Errorf("Failed to get next location. error:%v", err)
	}

	for _, loc := range resp.Results {
		fmt.Printf("---%s\n", loc.Name)
	}

	conf.NextURL = resp.Next
	conf.PrevURL = resp.Previous

	return nil
}

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Not enough args passed in.")
	}

	locName := args[0]

	resp, err := conf.client.GetAreaInfo(locName, cacheInterval)
	if err != nil {
		fmt.Errorf("Failed to get next location. error:%v", err)
	}

	fmt.Printf("---Pokemon found in: %s\n", resp.Name)
	for _, poke := range resp.PokemonEncounters {

		fmt.Printf("---%s\n", poke.Pokemon.Name)
	}

	return nil
}

func commandCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Not enough args passed in.")
	}

	pokeName := args[0]

	resp, err := conf.client.GetPokemon(pokeName, cacheInterval)
	if err != nil {
		fmt.Errorf("Failed to get next location. error:%v", err)
	}

	fmt.Printf("---Catching: %s\n", resp.Name)
	//the higher the base experience the higher to catch
	for resp.BaseExperience > rand.Int() {

	}

	return nil
}
