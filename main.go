package main

import (
	"time"

	"github.com/amalazimroshan/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute),
	}
	startRepl(&cfg)
}
