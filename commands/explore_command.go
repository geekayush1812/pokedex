package commands

import (
	"fmt"

	"github.com/geekayush1812/pokedex/pokeapi"
)

const pokeBaseUrl = "https://pokeapi.co/api/v2/location-area/"

func exploreCommand(cfg *ReplConfig, locationName string) error {
	response, err := pokeapi.GetPokemonNames(&cfg.HttpClient, pokeBaseUrl+locationName)

	if err != nil {
		return err
	}

	for _, pokemon := range response.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
