package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	var res []string
	for _, word := range words {
		if word != "" {
			word = strings.ToLower(word)
			res = append(res, word)
		}
	}
	return res
}

func startRepl() {
	initCommands()
	re := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		re.Scan()

		words := cleanInput(re.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]

		if val, exists := commands[command]; exists {
			err := val.callback(words)
			if err != nil{
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
