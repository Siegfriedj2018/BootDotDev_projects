package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func CommandPokedex(_ *Config, _ *internal.Cache, pokes *pokedex.Pokedex, _ ...string) error {
	pooks, err := pokes.GetAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pook := range pooks {
		fmt.Printf(" - %s\n", pook)
	}

	return nil
}