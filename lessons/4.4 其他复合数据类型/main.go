package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const RepositoriesURL = "https://api.github.com/search/repositories"

type RepositoriesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Repository
}

type Repository struct {
	Name      string
	CreatedAt time.Time `json:"created_at"`
	HtmlURL   string    `json:"html_url"`
	Owner     *Owner
}
type Owner struct {
	Login   string
	HtmlURL string `json:"html_url"`
}

func SearchRepositories(terms []string) (*RepositoriesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(RepositoriesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result RepositoriesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, err

}

func main() {
	result, err := SearchRepositories([]string{"vue"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d repositories: \n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("%-15s %9.9s %.10s\n",
			item.Name, item.Owner.Login, item.CreatedAt)
	}

}
