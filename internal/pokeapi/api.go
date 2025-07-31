package internal

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ZacharyLozevski/pokedexcli/config"
)

func makeAPIRequest(method, url string, body io.Reader) ([]byte, error) {
	// build the request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// make our client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read from the body data stream
	return io.ReadAll(res.Body)
}

type PokemonData struct {
	Next interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func LocationAreaAPIRequest(config *config.Config, url string) ([]string, error) {
	// make the api request and extract the body
	body, err := makeAPIRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// parse the body into a PokemonData struct
	var data PokemonData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// extract names
	var names []string
	for _, result := range data.Results {
		names = append(names, result.Name)
	}

	// update config
	if prev, ok := data.Previous.(string); ok {
		config.Previous = prev
	} else {
		config.Previous = ""
	}

	if next, ok := data.Next.(string); ok {
		config.Next = next
	} else {
		config.Next = ""
	}

	return names, nil
}