package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}

	//Create a new http get request, if error return the error
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	//Execute the http request, if error return error
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}

	//Read the response body into a variable
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	//Unmarshal the json data into a struct
	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	//Add the location data into the cache
	c.cache.Add(url, dat)
	return pokemonResp, nil
}
