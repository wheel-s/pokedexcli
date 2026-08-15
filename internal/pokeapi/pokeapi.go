package pokeapi

import "github.com/wheel-s/pokedexcli/internal/pokecache"

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache	pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client {
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type LocationAreasResp struct {
	Count	int	`json:"count"`
	Next	*string	`json:"next"`
	Previous	*string `json:"previous"`
	Results	[]struct {
		Name	string	`json:"name"`
		URL	string	`json:"url"`
	} `json:"results"`
}

type LocationArea struct {
	EncounterMethodRates	[]struct {
		EncounterMethod		struct {
			Name	string
			URL	string
		}
		VersionDetails	[]struct {
			Rate		int
			Version	[]struct {
				Name	string
				URL	string
			}
		}
	}

	GameIndex	int
	ID		int
	Location	struct {
		Name	string
		URL 	string
	}
	Name		string
	Names		[]struct {
		Language	struct {
			Name	string
			URL	string
		}
		Name	string
	}
	PokemonEncounters	[]struct {
		Pokemon		struct {
			Name	string
			URL	string
		}
		
		VersionDetails		[]struct {
			EncounterDetails	[]struct {
				Chance		int
				ConditionValues	[]interface{}
				MaxLevel	int
				Method		struct {
					Name	string
					URL 	string
				}
				MinLevel	int
			}
			MaxChance	int
			Version		struct {
				Name	string
				URL	string
			}
		}
	}	
}
