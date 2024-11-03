package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocationArea(location string) (ExploreLocationAreaResponse, error) {

	

	fullURL := baseURL + "/location-area/" + location

	if cacheData, ok := c.cache.Get(fullURL); ok {
		resp := ExploreLocationAreaResponse{}
		err := json.Unmarshal(cacheData, &resp)
		if err != nil {
			return ExploreLocationAreaResponse{}, err
		}

		return resp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ExploreLocationAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreLocationAreaResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return ExploreLocationAreaResponse{}, fmt.Errorf("Bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreLocationAreaResponse{}, err
	}

	c.cache.Add(fullURL, data)

	resp := ExploreLocationAreaResponse{}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return ExploreLocationAreaResponse{}, err
	}

	return resp, nil
}
