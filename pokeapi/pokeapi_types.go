package pokeapi

import (
	"net/http"

	"github.com/geekayush1812/pokedex/pokecache"
)

type result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeLocationApiResponse struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []result `json:"results"`
}

type PokeNameApiResponse struct {
	PokemonEncounters []struct {
		Pokemon result `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonStats struct {
	BaseExperience         int    `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Name                   string `json:"name"`
	Stats                  []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type ApiResponse interface {
	PokeLocationApiResponse | PokeNameApiResponse | PokemonStats
}

type Client struct {
	Cache      pokecache.Cache
	httpClient http.Client
}
