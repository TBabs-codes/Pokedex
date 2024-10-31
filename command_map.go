package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	next     *string
	previous *string
}

func callbackMap(nextPrev config) error {
	fmt.Println("Start of callback Map")
	url := "https://pokeapi.co/api/v2/location-area/"
	fmt.Println(nextPrev.next)
	if nextPrev.next != nil {
		url = *nextPrev.next
	}
	//send request, retrieve response
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred during http GET request.")
		return err
	}
	defer resp.Body.Close()

	//decode response
	locationResp := LocationAreaResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationResp); err != nil {
		fmt.Println("Error occurred decoding.")
		return err
	}
	//Print response
	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}

	nextPrev.next = locationResp.Next
	nextPrev.previous = locationResp.Previous


	fmt.Println(*nextPrev.next)
	fmt.Println(nextPrev.previous)
	return nil
}

func callbackMapB(nextPrev config) error {
	fmt.Println("Start of callback Map")
	url := "https://pokeapi.co/api/v2/location-area/"
	
	if nextPrev.previous != nil {
		url = *nextPrev.previous
	}else {
		fmt.Println("Error: This is the top of the location areas list.")
	}
	//send request, retrieve response
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred during http GET request.")
		return err
	}
	defer resp.Body.Close()

	//decode response
	locationResp := LocationAreaResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationResp); err != nil {
		fmt.Println("Error occurred decoding.")
		return err
	}
	//Print response
	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}

	nextPrev.next = locationResp.Next
	nextPrev.previous = locationResp.Previous


	fmt.Println(*nextPrev.next)
	fmt.Println(nextPrev.previous)
	return nil
}