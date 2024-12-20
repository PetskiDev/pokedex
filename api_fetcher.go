package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2"


func getJsonFrom(path string) ([]byte, error){
	fullPath := baseUrl + path
	req, err := http.NewRequest("GET", fullPath, nil)
	
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	return io.ReadAll(res.Body)
}


func getLocation(offset, limit int) ([]Location, error){
	query := fmt.Sprintf("/location-area?offset=%v&limit=%v", offset, limit)
	data, err := getJsonFrom(query)
	if err != nil{
		return nil, err
	}
	var rawResponse LocationResponse
	err = json.Unmarshal(data, &rawResponse)
	if err != nil{
		return nil, err
	}
	return rawResponse.Results, nil
}