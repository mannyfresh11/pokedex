package main

import (
	"fmt"
	"os"
	"time"
)

var cacheInterval = time.Hour

func commandHelp(cong *config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	cmd := Commands()

	for _, c := range cmd {
		fmt.Printf("- %s: %s\n", c.Name, c.Description)
	}

	fmt.Println("")

	return nil
}

func commandExit(conf *config) error {

	os.Exit(0)

	return nil
}

func commandMap(conf *config) error {

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

func commandMapb(conf *config) error {

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
