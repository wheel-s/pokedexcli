package main

import "github.com/wheel-s/pokedexcli/internal/pokeapi"

import (
	"fmt"
	"log"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()
	
	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location area's:")
	for _, area := range resp.Results{
		fmt.Printf(" - %s\n", area.Name)
	}
	return nil
} 

