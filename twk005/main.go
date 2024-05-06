package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Country represents the relevant fields of each country received from the API
type Country struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
}

type APIResponse struct {
	Page       int       `json:"page"`
	TotalPages int       `json:"total_pages"`
	Data       []Country `json:"data"`
}

func main() {
	substring := "in"        // The substring to search in country names
	minPopulation := 1000000 // The minimum population threshold

	count, err := getCountries2(substring, minPopulation)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Number of countries: %d\n", count)
	}
}

func getCountries2(s string, p int) int {
	count := 0
	page := 1

	for {
		url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/countries/search?name=%s&page=%d", s, page)
		response, err := http.Get(url)
		if err != nil {
			return 0
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return 0
		}

		var apiResponse APIResponse
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			return 0
		}

		for _, country := range apiResponse.Data {
			if strings.Contains(strings.ToLower(country.Name), strings.ToLower(s)) && country.Population >= p {
				count++
			}
		}

		if page >= apiResponse.TotalPages {
			break
		}
		page++
	}

	return count
}

func getCountries(substring string, minPopulation int) (int, error) {
	url := "https://jsonmock.hackerrank.com/api/countries/search?name=" + substring
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, country := range apiResponse.Data {
		if strings.Contains(strings.ToLower(country.Name), strings.ToLower(substring)) && country.Population >= minPopulation {
			count++
		}
	}

	return count, nil
}
