package commands

import "fmt"

func pokedexCommand(_ *ReplConfig, _ string) error {

	if len(caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range caughtPokemon {
		fmt.Printf("\t-%s\n", pokemon.Name)
	}

	return nil
}
