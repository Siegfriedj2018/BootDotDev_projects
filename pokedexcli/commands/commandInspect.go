package commands

import (
	"fmt"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func CommandInspect(conf *Config, cash *internal.Cache, pd *pokedex.Pokedex, pokeMon ...string) error {
	if len(pokeMon) == 0 {
		return fmt.Errorf("please enter a pokemon name, type help for usage")
	}
	url := internal.BaseURL + "pokemon/" + pokeMon[0]
	data, ok := pd.Get(url)
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", data.Name, data.Height, data.Weight)
	for _, stat := range data.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ty := range data.Types {
		fmt.Printf("  - %s\n", ty.Type.Name)
	}
	return nil
}