package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(areaName string) (RespDeepLocation, error) {
	//If no area name provided, default to first area in the list
	if areaName == "" {
		areaName = "canalave-city-area"
	}

	//Else construct the URL from the given area name
	url := baseURL + "/location-area/" + areaName

	//Check if a cached version of the data exists and use that if so
	if val, ok := c.cache.Get(url); ok {
		locationResp := RespDeepLocation{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespDeepLocation{}, err
		}
		return locationResp, nil
	}

	//Create a new http get request, if error return the error
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespDeepLocation{}, err
	}

	//Execute the http request, if error return error
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespDeepLocation{}, err
	}
	defer resp.Body.Close() //Close the http response when the function exits

	//Read the response body into a variable
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespDeepLocation{}, err
	}

	//Unmarshal the json data into a struct
	locationResp := RespDeepLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespDeepLocation{}, err
	}

	//Add the location data into the cache
	c.cache.Add(url, dat)
	return locationResp, nil
}
