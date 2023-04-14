package commands

import (
	"fmt"
	"math/rand"

	"github.com/geekayush1812/pokedex/pokeapi"
)

const pokemonStatsUrl = "https://pokeapi.co/api/v2/pokemon/"

var caughtPokemon = make(map[string]pokeapi.PokemonStats)

func catchCommand(cfg *ReplConfig, pokemonName string) error {

	if _, ok := caughtPokemon[pokemonName]; ok {
		fmt.Printf("%s is already caught!\n", pokemonName)
		return nil
	}

	pokemon, err := pokeapi.GetPokemonStats(&cfg.HttpClient, pokemonStatsUrl+pokemonName)

	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	caughtPokemon[pokemon.Name] = pokemon
	return nil
}
