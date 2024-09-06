package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/mannyfresh11/pokedex/internal/pokicache"
)

var cacheInterval = time.Hour

func commandHelp(cong *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("No need to pass argument")
	}

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
	if len(args) > 0 {
		return errors.New("No need to pass argument")
	}

	os.Exit(0)

	return nil
}

func commandMap(conf *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("No need to pass argument")
	}

	resp, err := conf.client.GetLocation(conf.nextURL, cacheInterval)
	if err != nil {
		return fmt.Errorf("Failed to get next location. error:%v", err)
	}

	for _, loc := range resp.Results {
		fmt.Printf("---%s\n", loc.Name)
	}

	conf.nextURL = resp.Next
	conf.prevURL = resp.Previous

	return nil
}

func commandMapb(conf *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("No need to pass argument")
	}

	resp, err := conf.client.GetLocation(conf.prevURL, cacheInterval)
	if err != nil {
		return fmt.Errorf("Failed to get next location. error:%v", err)
	}

	for _, loc := range resp.Results {
		fmt.Printf("---%s\n", loc.Name)
	}

	conf.nextURL = resp.Next
	conf.prevURL = resp.Previous

	return nil
}

func commandExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Not enough args passed in.")
	}

	locName := args[0]

	resp, err := conf.client.GetAreaInfo(locName, cacheInterval)
	if err != nil {
		return fmt.Errorf("Failed to get next location. error:%v", err)
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

	pokemon, err := conf.client.GetPokemon(pokeName, cacheInterval)
	if err != nil {
		return fmt.Errorf("Failed to get pokemon. - %v\n", err)
	}

	fmt.Printf("---Catching: %s\n", pokemon.Name)

	rand := rand.Intn(pokemon.BaseExperience)
	exp := pokemon.BaseExperience
	var div float32

	switch {
	case exp >= 300:
		div = .9
	case exp < 300 && exp >= 200:
		div = .7
	case exp < 200 && exp >= 100:
		div = .5
	default:
		div = .3
	}

	if float32(rand) > float32(exp)*div {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		conf.pokemon.AddPokemon(pokemon.Name, pokicache.Pokemon(pokemon))
	} else {
		return fmt.Errorf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(conf *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("Not enough args passed in.")
	}

	pokeName := args[0]

	pokemon, ok := conf.pokemon.GetPokemon(pokeName)
	if !ok {
		return fmt.Errorf("No pokemon %s caught yet!\n", pokeName)
	}

	fmt.Printf("- Name: %s\n", pokemon.Name)
	fmt.Printf("- Height: %d\n", pokemon.Height)
	fmt.Printf("- Weight: %d\n", pokemon.Weight)
	fmt.Println("- Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("--- %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("-Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("--- %s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(conf *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("No need to pass argument")
	}

	pokemons := conf.pokemon.GetAllPokemon()
	if len(pokemons) == 0 {
		return fmt.Errorf("No pokemon caught yet!\n")
	}

	for pokemon, _ := range pokemons {
		fmt.Printf("--- %s\n", pokemon)
	}

	return nil
}
