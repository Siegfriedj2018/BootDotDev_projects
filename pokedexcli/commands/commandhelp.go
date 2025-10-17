package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func CommandHelp(conf *Config, cache *internal.Cache, _ *pokedex.Pokedex, extra ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range GetCommands() {
		fmt.Printf("\t%v: %v\n", cmd.Name, cmd.Description)
	}
	return nil
}