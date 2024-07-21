package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonList(loc string) (jsonPokemon, error) {
	url := baseURL + "location-area/" + loc

	if cacheData, ok := c.cache.Get(url); ok {
		locationResp := jsonPokemon{}
		err := json.Unmarshal(cacheData, &locationResp)
		if err != nil {
			return jsonPokemon{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return jsonPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return jsonPokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return jsonPokemon{}, err
	}

	c.cache.Add(url, data)

	pokemonResp := jsonPokemon{}
	json.Unmarshal(data, &pokemonResp)

	return pokemonResp, nil
}
