package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/geekayush1812/pokedex/pokecache"
)

func NewClient(cacheReapInterval time.Duration) Client {
	return Client{
		Cache:      pokecache.NewCache(cacheReapInterval),
		httpClient: http.Client{},
	}
}

func getApiResponse[T ApiResponse](c *Client, url string) (T, error) {

	var res T

	if byteResponse, hasCache := c.Cache.Get(url); hasCache {
		err := json.Unmarshal(byteResponse, &res)
		return res, err
	}

	apiResponse, err := c.httpClient.Get(url)

	if err != nil {
		return res, err
	}

	defer apiResponse.Body.Close()

	err = json.NewDecoder(apiResponse.Body).Decode(&res)

	if err != nil {
		return res, err
	}

	marshalledApiResponse, err := json.Marshal(res)

	if err != nil {
		return res, err
	}

	c.Cache.Add(url, marshalledApiResponse)

	return res, nil
}

func GetLocationAreas(c *Client, url string) (PokeLocationApiResponse, error) {
	return getApiResponse[PokeLocationApiResponse](c, url)
}

func GetPokemonNames(c *Client, url string) (PokeNameApiResponse, error) {
	return getApiResponse[PokeNameApiResponse](c, url)
}

func GetPokemonStats(c *Client, url string) (PokemonStats, error) {
	return getApiResponse[PokemonStats](c, url)
}
