package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
	"os"
)

func CommandExit(conf *Config, cache *internal.Cache, _ *pokedex.Pokedex, extra ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
