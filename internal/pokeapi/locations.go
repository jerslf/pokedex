package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ListLocations() (LocationAreasResp, error) {
	fullURL := baseURL + "/location-area"

	res, err := http.Get(fullURL)
	if err != nil {
		return LocationAreasResp{}, err
	}
	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(body, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
