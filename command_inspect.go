package main


import (
	"errors"
	"fmt"
)


func callbackInspect(cfg *config, args ...string) error {
	if len(args) !=1 {
		return errors.New("no pokemon area provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)	
	fmt.Printf("height: %v\n", pokemon.Height)
	fmt.Printf("weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v", stat.Stat.Name, stat.BaseStat)
	}
	for _, typ := range pokemon.Types{
		fmt.Printf(" - %s", typ.Type.Name)
	}
	return nil
}
