package main

import (
	"time"

	"github.com/mannyfresh11/pokedex/internal/api"
)

type config struct {
	client   api.Client
	NextURL  *string
	PrevURL  *string
	areaName *string
}

func main() {

	cacheInterval := time.Hour
	conf := config{
		client: api.NewClient(cacheInterval),
	}

	startREPL(&conf)
}
