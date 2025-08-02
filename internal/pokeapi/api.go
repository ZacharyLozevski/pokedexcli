package pokeapi

import (
	"io"
	"net/http"
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

func locationAreaAPIRequest(url string) ([]byte, error) {
	// make an api request
	rawData, err := makeAPIRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}