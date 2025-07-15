package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("pokemon name not provided")
	}

	pokemonData, err := cfg.pokeapiClient.GetPokemonData(args[0])

	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())

	baseExp := pokemonData.BaseExperience
	difficulty := int(baseExp / 10)
	pokemonName := pokemonData.Name

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	roll := rand.Intn(100-1) + 1
	if roll > difficulty {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
	}
	return nil
}
