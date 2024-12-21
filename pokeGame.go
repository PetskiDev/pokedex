package main

import (
	"fmt"
	"time"

	"github.com/PetarKovacovski/pokedex/internal/pokecache"
)

var state = struct {
	lastPrinted   int
	limit         int
	cache         pokecache.Cache
	pokemonCaught map[string]Pokemon
}{
	lastPrinted:   -MAP_EXPLORE, //NEEDS TO BE SAME AS LIMIT
	limit:         MAP_EXPLORE,
	cache:         pokecache.NewCache(CACHE_IN_S * time.Second),
	pokemonCaught: make(map[string]Pokemon),
}

func getNextLocations() ([]Location, error) {
	state.lastPrinted += state.limit
	locations, err := getLocation(state.lastPrinted, state.limit)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func getPrevLocations() ([]Location, error) {
	if state.lastPrinted <= 0 {
		return nil, fmt.Errorf("can't go back")
	}
	state.lastPrinted -= state.limit
	locations, err := getLocation(state.lastPrinted, state.limit)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (p *Pokemon) printPokemon() {
	fmt.Println("Height:", p.Height)
	fmt.Println("Weight:", p.Weight)

	fmt.Println("Stats:")
	for _, val := range p.Stats {
		fmt.Printf("\t-%v: %v\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types:")
	for _, val := range p.Types {
		fmt.Printf("\t-%v\n", val.Type.Name)
	}

}
