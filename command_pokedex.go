package main

import (
    "fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
    fmt.Println("Pokemon in Pokedex:")
    for _, pokemon := range cfg.caughtPokemon {
        fmt.Println(" - %s\n", pokemon.Name)
    }

    return nil
}
