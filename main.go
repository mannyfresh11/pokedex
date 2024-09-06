package main

import (
	"time"

	"github.com/mannyfresh11/pokedex/internal/api"
	"github.com/mannyfresh11/pokedex/internal/pokicache"
)

type config struct {
	client  api.Client
	pokemon pokicache.PokemonData
	nextURL *string
	prevURL *string
}

func main() {

	cacheInterval := time.Hour
	conf := config{
		client:  api.NewClient(cacheInterval),
		pokemon: pokicache.NewPokemonCache(),
	}

	startREPL(&conf)
}
