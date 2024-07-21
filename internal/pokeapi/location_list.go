package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageURL *string) (jsonLocations, error) {
	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if cacheData, ok := c.cache.Get(url); ok {
		locationResp := jsonLocations{}
		err := json.Unmarshal(cacheData, &locationResp)
		if err != nil {
			return jsonLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return jsonLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return jsonLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return jsonLocations{}, err
	}

	c.cache.Add(url, data)

	locationResp := jsonLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return jsonLocations{}, err
	}

	return locationResp, nil
}
