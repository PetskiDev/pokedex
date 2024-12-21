package main

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous any        `json:"previous"`
	Results  []Location `json:"results"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Encounter struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Encounter `json:"pokemon"`
}

type PokemonEncounters struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
