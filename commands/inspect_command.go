package commands

import (
	"errors"
	"fmt"
)

func inspectCommand(cfg *ReplConfig, pokemonName string) error {

	pokemonStats, hasPokemonBeenCaught := caughtPokemon[pokemonName]

	if !hasPokemonBeenCaught {
		return errors.New("pokemon has not been caught")
	}

	fmt.Println("Name: ", pokemonStats.Name)
	fmt.Println("Height: ", pokemonStats.Height)
	fmt.Println("Weight: ", pokemonStats.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonStats.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemonStats.Types {
		fmt.Printf("\t-%s\n", pokeType.Type.Name)
	}

	return nil
}
