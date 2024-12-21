package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2"

func getJsonFrom(path string) ([]byte, error) {
	fullPath := baseUrl + path

	if c, has := state.cache.Get(fullPath); has {
		return c, nil
	}
	req, err := http.NewRequest("GET", fullPath, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	toReturn, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	state.cache.Add(fullPath, toReturn)
	return toReturn, nil
}

func getLocation(offset, limit int) ([]Location, error) {
	query := fmt.Sprintf("/location-area?offset=%v&limit=%v", offset, limit)
	data, err := getJsonFrom(query)
	if err != nil {
		return nil, err
	}
	var rawResponse LocationResponse
	err = json.Unmarshal(data, &rawResponse)
	if err != nil {
		return nil, err
	}
	return rawResponse.Results, nil
}

func findPokemonInMap(mapName string) ([]PokemonEncounter, error) {
	query := fmt.Sprintf("/location-area/%v", mapName)
	data, err := getJsonFrom(query)
	if err != nil {
		return nil, err
	}
	var unmarshalData PokemonEncounters
	err = json.Unmarshal(data, &unmarshalData)
	if err != nil {
		return nil, err
	}
	return unmarshalData.PokemonEncounters, nil
}

func getPokemonInfo(pokemonName string) (Pokemon, error) {
	query :=  fmt.Sprintf("/pokemon/%v", pokemonName)

	data, err := getJsonFrom(query)
	
	if err != nil {
		return Pokemon{}, err
	}
	
	var unmarshalData Pokemon
	err = json.Unmarshal(data, &unmarshalData)
	if err != nil {
		return Pokemon{}, err
	}
	
	return unmarshalData, nil
}
