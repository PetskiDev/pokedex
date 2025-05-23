package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

var commands map[string]cliCommand

func initCommands() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "show the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "show the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "list pokemon found in location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a given pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Print stats of a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemon caught",
			callback:    commandPokedex,
		},
	}
}

func commandMap(args []string) error {
	res, err := getNextLocations()
	if err != nil {
		return err
	}
	for _, loc := range res {
		fmt.Println(loc.Name)
	}
	return nil
}
func commandMapB(args []string) error {
	res, err := getPrevLocations()
	if err != nil {
		return err
	}
	for _, loc := range res {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(args []string) error {
	fmt.Println("Welcome to the Pokedex!")

	fmt.Println("Usage:")
	fmt.Println()

	for key, val := range commands {
		fmt.Printf("%v: %v\n", key, val.description)
	}
	return nil
}

func commandExplore(args []string) error {
	if len(args) <= 1 {
		return errors.New("you must provide a location name")
	}

	mapName := args[1]
	fmt.Printf("Exploring %v...\n", mapName)

	allEncounter, err := findPokemonInMap(mapName)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")

	for _, enc := range allEncounter {
		fmt.Println(" - " + enc.Name)
	}
	return nil
}

func commandCatch(args []string) error {
	if len(args) <= 1 {
		return errors.New("you must provide a location name")
	}

	pokemonName := args[1]

	if _, in := state.pokemonCaught[pokemonName]; in {
		return fmt.Errorf("%v is already caught", pokemonName)
	}

	pokemon, err := getPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v...", pokemonName)

	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	state.pokemonCaught[pokemonName] = pokemon

	return nil
}

func commandInspect(args []string) error {
	if len(args) <= 1 {
		return errors.New("you must provide a location name")
	}

	pokemonName := args[1]

	pokemon, in := state.pokemonCaught[pokemonName]
	if !in {
		return fmt.Errorf("you have not caught that pokemon")
	}

	pokemon.printPokemon()
	return nil
}

func commandPokedex(args []string) error {
	if len(state.pokemonCaught) == 0 {
		return fmt.Errorf("you don't have any pokemon")
	}
	fmt.Println("Your Pokedex:")
	for name := range state.pokemonCaught {
		fmt.Printf("\t- %v\n", name)
	}
	return nil
}
