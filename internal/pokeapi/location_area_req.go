package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(nextURL *string) (LocationAreaResponse, error) {

	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if nextURL != nil {
		fullURL = *nextURL
	}

	if cacheData, ok := c.cache.Get(fullURL); ok {
		locAreaResponse := LocationAreaResponse{}
		err := json.Unmarshal(cacheData, &locAreaResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locAreaResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullURL, data)

	locAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locAreaResponse, nil
}
