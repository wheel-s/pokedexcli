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

