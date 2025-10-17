package commands

import (
	"fmt"
	"math/rand"
	"go-tutorial/bootdotdev/pokedexcli/internal"
	"go-tutorial/bootdotdev/pokedexcli/pokedex"
)

func CommandCatch(conf *Config, cash *internal.Cache, pooDex *pokedex.Pokedex, pokemon ...string) error {
	if len(pokemon) == 0 {
		return fmt.Errorf("please enter a pokemon to catch, type help for usage")
	}
	url := internal.BaseURL + "pokemon/" + pokemon[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon[0])
	res, err := internal.CatchPokemon(url, cash)
	if err != nil {
		return fmt.Errorf("there was an internal error: %w", err)
	}

	catchChance := clamp(res.BaseExperience)
	randCatch := rand.Float64()
	fmt.Printf("Catch Chace is: %.2f, random Catch: %.2f\n", catchChance, randCatch)

	if randCatch <= catchChance {
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemon[0])
		pooDex.Add(url, *res)
	} else {
		fmt.Printf("%s escaped!\n", pokemon[0])
	}

	return nil
}

func clamp(xp int) float64 {
	ceiling := 0.9
	floor := 0.1
	base_xp_mid := 300
	k := (ceiling - 0.5) / float64(base_xp_mid)
	
	chance := ceiling - (float64(xp) * k)
	if chance >= ceiling {
		return ceiling
	} else if chance <= floor {
		return floor
	}
	return chance
}