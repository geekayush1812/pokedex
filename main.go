package main

import (
	"time"

	"github.com/geekayush1812/pokedex/commands"
	"github.com/geekayush1812/pokedex/pokeapi"
)

func main() {
	cliRepl := commands.ReplConfig{
		HttpClient: pokeapi.NewClient(2 * time.Minute),
	}

	cliRepl.Start()
}
