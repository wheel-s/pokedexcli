package main

import "fmt"

func callbackHelp() error{
	fmt.Println("Welcome to rhe Pokedex help menu")
	fmt.Println("Here are ypur available commands")
	fmt.Println(" - help")
	fmt.Println(" - exit")	
	fmt.Println(" - map")
	fmt.Println("")

	return nil
}
