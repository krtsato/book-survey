package services

import (
	"encoding/json"
	"fmt"
)

// by JSON-to-GO (https://mholt.github.io/json-to-go/)
type DecResBodyType struct {
	Items []struct {
		Title         string `json:"title"`
		SalesDate     string `json:"salesDate"`
		ItemPrice     int    `json:"itemPrice"`
		BooksGenreID  string `json:"booksGenreId"`
		ReviewCount   int    `json:"reviewCount"`
		ReviewAverage string `json:"reviewAverage"`
		ItemURL       string `json:"itemUrl"`
	} `json:"Items"`
	PageCount        int `json:"pageCount"`
	Hits             int `json:"hits"`
	First            int `json:"first"`
	Last             int `json:"last"`
	Count            int `json:"count"`
	Page             int `json:"page"`
	GenreInformation []struct {
		Current []struct {
			BooksGenreID   string `json:"booksGenreId"`
			BooksGenreName string `json:"booksGenreName"`
		} `json:"current"`
		Children []struct {
			BooksGenreID   string `json:"booksGenreId"`
			BooksGenreName string `json:"booksGenreName"`
		} `json:"children"`
	} `json:"GenreInformation"`
}

func DecordResponse(resBody []byte) (*DecResBodyType, error) {
	decResBody := &DecResBodyType{}
	if err := json.Unmarshal(resBody, decResBody); err != nil {
		return &DecResBodyType{}, fmt.Errorf("While unmarshalling the response json: %w", err)
	}
	return decResBody, nil
}
