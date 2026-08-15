package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error{
	fmt.Println("Welcome to rhe Pokedex help menu")
	fmt.Println("Here are ypur available commands")
	
	availableCommands := getCommands()
	for _, cmd := range availableCommands{
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")

	return nil
}
