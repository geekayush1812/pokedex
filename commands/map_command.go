package commands

import (
	"errors"
	"fmt"

	"github.com/geekayush1812/pokedex/pokeapi"
)

const numberOfLocationAreas = 20

type MapCommandConfig struct {
	NextUrl *string
	PrevUrl *string
}

var locationNextUrl = "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"

var mapCommandConfig = MapCommandConfig{
	NextUrl: &locationNextUrl,
	PrevUrl: nil,
}

func mapCommand(cfg *ReplConfig, _ string) error {
	return listLocations(cfg, &mapCommandConfig, *mapCommandConfig.NextUrl)
}

func mapbCommand(cfg *ReplConfig, _ string) error {
	if mapCommandConfig.PrevUrl == nil {
		return errors.New("no previous location areas")
	}

	return listLocations(cfg, &mapCommandConfig, *mapCommandConfig.PrevUrl)
}

func listLocations(cfg *ReplConfig, mcfg *MapCommandConfig, url string) error {

	response, err := pokeapi.GetLocationAreas(&cfg.HttpClient, url)

	if err != nil {
		return err
	}

	results := response.Results

	for i := 0; i < numberOfLocationAreas; i++ {
		fmt.Println(results[i].Name)
	}

	mcfg.NextUrl = response.Next
	mcfg.PrevUrl = response.Previous

	return nil
}
