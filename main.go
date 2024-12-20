package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string{
	words := strings.Split(text, " ")
	var res []string
	for _, word := range words{
		if word != ""{
			word = strings.ToLower(word)
			res = append(res, word)
		}
	}
	return res
}

func main(){

	for {
		fmt.Print("Pokedex > ")
		re := bufio.NewScanner(os.Stdin)
		re.Scan()
		command := cleanInput(re.Text())
		fmt.Printf("Your command was: %v\n", command[0])
	}

	
}