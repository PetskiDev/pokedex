package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
		"map":{
			name:        "map",
			description: "show the next 20 locations",
			callback:    commandMap,
		},
		"mapb":{
			name:        "mapb",
			description: "show the previous 20 locations",
			callback:    commandMapB,
		},
	}
}

func commandMap() error{
	res, err := getNextLocations()
	if err != nil{
		return err
	}
	for _, loc := range res{
		fmt.Println(loc.Name)
	}
	return nil
}
func commandMapB() error{
	res, err := getPrevLocations()
	if err != nil{
		return err
	}
	for _, loc := range res{
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")

	fmt.Println("Usage:")
	fmt.Println()

	for key, val := range commands {
		fmt.Printf("%v: %v\n", key, val.description)
	}
	return nil
}

