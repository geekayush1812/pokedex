package commands

import (
	"errors"
	"os"
)

func exit(cfg *ReplConfig, _ string) error {
	os.Exit(0)
	return errors.New("Exiting Pokedex CLI")
}
