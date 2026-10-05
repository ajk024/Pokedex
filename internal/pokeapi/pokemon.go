package pokeapi

type Pokedex map[string]Pokemon

type Pokemon struct {
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats  []struct {
		//BaseStat int `json:"base_stat"`
		//Effort   int `json:"effort"`
		Stat struct {
			Name string `json:"name"`
			//URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func NewPokedex() Pokedex {
	return Pokedex{}
}

type PokemonData struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type EncounterData struct {
	//ID   int    `json:"id"`
	Name string `json:"name"`
	/*
		GameIndex            int    `json:"game_index"`
		EncounterMethodRates []struct {
			EncounterMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"encounter_method"`
			VersionDetails []struct {
				Rate    int `json:"rate"`
				Version struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"version"`
			} `json:"version_details"`
		} `json:"encounter_method_rates"`
		Location struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"location"`
		Names []struct {
			Name     string `json:"name"`
			Language struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"names"`
	*/
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		/*
			VersionDetails []struct {
				Version struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"version"`
				MaxChance        int `json:"max_chance"`
				EncounterDetails []struct {
					MinLevel int `json:"min_level"`
					MaxLevel int `json:"max_level"`
					Chance   int `json:"chance"`
					Method   struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					} `json:"method"`
					ConditionValues []any `json:"condition_values"`
					PokemonDetails  any   `json:"pokemon_details"`
				} `json:"encounter_details"`
			} `json:"version_details"`
		*/
	} `json:"pokemon_encounters"`
}
