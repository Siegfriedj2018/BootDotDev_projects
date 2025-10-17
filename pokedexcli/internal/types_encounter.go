package internal

type EncounterMethods struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethodRates struct {
	EncounterMethods EncounterMethods 	`json:"encounter_method"`
	EncounterVersion []EncounterVersion `json:"version_details"`
}

type EncounterMethod struct {
	Name string	 `json:"name"`
	Url  *string `json:"url"`
}

type EncounterVersions struct {
	Rate						 int 							`json:"rate"`
	EncounterVersion EncounterVersion `json:"version_details"`
}

type EncounterVersion struct {
	Rate 		int 		`json:"rate"`
	Version Version `json:"version"`
}

type Version struct {
	Name string 	`json:"name"`
	Url  *string 	`json:"url"`
}

type Location struct {
	Name string  `json:"name"`
	Url  *string `json:"url"`
}

type Names struct {
	Name    	*string  `json:"name"`
	Language	Language `json:"language"`
}

type Language struct {
	Name  string	`json:"name"`
	Url   *string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon				 Pokemon 			  	`json:"pokemon"`
	PokemonVersion []PokemonVersion `json:"version_details"`
}

type Pokemon struct {
	Name	string		`json:"name"`
	Url		*string		`json:"url"`
}

type PokemonVersion struct {
	Version						Version							`json:"version"`
	MaxChance					int									`json:"max_chance"`
	EncounterDetails	[]EncounterDetails 	`json:"encounter_details"`
}

type EncounterDetails struct {
	MinLevel					int 			`json:"min_level"`
	MaxLevel					int				`json:"max_level"`
	Condition_values 	[]any 	 	`json:"condition_values"`
	Chance						int 			`json:"chance"`
	Method  					Method 		`json:"method"`
}

type Method struct {
	Name	string		`json:"name"`
	Url		*string		`json:"url"`
}