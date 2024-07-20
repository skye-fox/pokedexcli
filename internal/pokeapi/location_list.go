package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationList(pageURL *string) (JsonLocations, error) {
	url := baseURL + "location/"
	if pageURL != nil {
		url = *pageURL
	}

	if cacheData, ok := c.cache.Get(url); ok {
		locationResp := JsonLocations{}
		err := json.Unmarshal(cacheData, &locationResp)
		if err != nil {
			return JsonLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return JsonLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return JsonLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return JsonLocations{}, err
	}

	c.cache.Add(url, data)

	locationResp := JsonLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return JsonLocations{}, err
	}

	return locationResp, nil
}
