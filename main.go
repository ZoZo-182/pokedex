package main

import (
    "github.com/ZoZo-182/pokedexcli/internal/pokeapi"
    "time"
)

type config struct {
    pokeapiClient       pokeapi.Client
    nextLocationAreaURL *string
    prevLocationAreaURL *string
    caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
    cfg := config {
        pokeapiClient: pokeapi.NewClient(time.Hour),
        caughtPokemon: make(map[string]pokeapi.Pokemon),
    }



    startRepl(&cfg)
}
