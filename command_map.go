package main

import (
    "fmt"
//    "log"
    //"github.com/ZoZo-182/pokedexcli/internal/pokeapi"
    "errors"
)

func callbackMap(cfg *config, args ...string) error {
    resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
    if err != nil {
        return err
    }
    fmt.Println("Location areas:")
    for _, area := range resp.Results {
        fmt.Printf(" - %s\n", area.Name)
    }
    cfg.nextLocationAreaURL = resp.Next
    cfg.prevLocationAreaURL = resp.Previous

    return nil
}

func callbackMapb(cfg *config, args ...string) error {
    if cfg.prevLocationAreaURL == nil {
        return errors.New("you're on the first page")
    }

    resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
    if err != nil {
        return err
    }
    fmt.Println("Location areas:")
    for _, area := range resp.Results {
        fmt.Printf(" - %s\n", area.Name)
    }
    cfg.nextLocationAreaURL = resp.Next
    cfg.prevLocationAreaURL = resp.Previous

    return nil
}
