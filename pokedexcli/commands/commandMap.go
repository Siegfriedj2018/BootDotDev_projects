package commands

import (
	"fmt"

	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func CommandMap(conf *Config, cache *internal.Cache, _ *pokedex.Pokedex, extra ...string) error {
	url := internal.BaseURL + "location-area"
	if conf.Next != nil {
		url = *conf.Next
	} else {
		fmt.Println("You have reached the end of the map locations, starting over...")
	}
	
	response, err := internal.Connection(url, cache)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for idx, result := range response.Results {
		fmt.Printf("  %d. %s\n", idx+1, result.Name)
	}

	conf.Next = response.Next
	conf.Previous = response.Previous
	return nil
}

func CommandMapb(conf *Config, cache *internal.Cache, _ *pokedex.Pokedex, extra ...string) error {
	url := internal.BaseURL + "location-area"
	if conf.Previous != nil {
		url = *conf.Previous
	} else {
		fmt.Println("You have reached the beginning")
	}
	
	response, err := internal.Connection(url, cache)
	if err != nil {
		return err
	}
	fmt.Println("Locations:")
	for idx, result := range response.Results {
		fmt.Printf("  %d. %s\n", idx+1, result.Name)
	}

	conf.Next = response.Next
	conf.Previous = response.Previous
	return nil
}
