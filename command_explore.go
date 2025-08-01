package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("area name not provided")
	}

	locationResp, err := cfg.pokeapiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationResp.Location.Name)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
