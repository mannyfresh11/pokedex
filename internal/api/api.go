package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/mannyfresh11/pokedex/internal/pokicache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokicache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokicache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) GetLocation(url *string, cacheInterval time.Duration) (LocationArea, error) {

	endpoint := "/location-area"
	fullPath := baseURL + endpoint
	if url != nil {
		fullPath = *url
	}

	data, ok := c.cache.Get(fullPath)
	if ok {
		LA := LocationArea{}

		err := json.Unmarshal(data, &LA)
		if err != nil {
			return LocationArea{}, err
		}

		return LA, nil
	}

	resp, err := c.httpClient.Get(fullPath)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	LA := LocationArea{}

	err = json.Unmarshal(data, &LA)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullPath, data)

	return LA, nil
}

func (c *Client) GetAreaInfo(areaName string, cacheInterval time.Duration) (LocationAreaName, error) {

	endpoint := "/location-area/" + areaName
	fullPath := baseURL + endpoint

	data, ok := c.cache.Get(fullPath)
	if ok {
		LA := LocationAreaName{}

		err := json.Unmarshal(data, &LA)
		if err != nil {
			return LocationAreaName{}, err
		}

		return LA, nil
	}

	resp, err := c.httpClient.Get(fullPath)
	if err != nil {
		return LocationAreaName{}, err
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaName{}, err
	}

	LA := LocationAreaName{}

	err = json.Unmarshal(data, &LA)
	if err != nil {
		return LocationAreaName{}, err
	}

	c.cache.Add(fullPath, data)

	return LA, nil
}

func (c *Client) GetPokemon(pokemon string, cacheInterval time.Duration) (Pokemon, error) {

	endpoint := "/pokemon/" + pokemon
	fullPath := baseURL + endpoint

	data, ok := c.cache.Get(fullPath)
	if ok {
		pkm := Pokemon{}

		err := json.Unmarshal(data, &pkm)
		if err != nil {
			return Pokemon{}, err
		}

		return pkm, nil
	}

	resp, err := c.httpClient.Get(fullPath)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pkm := Pokemon{}

	err = json.Unmarshal(data, &pkm)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullPath, data)

	return pkm, nil
}
