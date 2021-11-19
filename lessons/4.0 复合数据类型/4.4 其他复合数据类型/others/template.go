package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const RepositoriesURL = "https://api.github.com/search/repositories"

type RepositoriesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Repository
}
type Owner struct {
	Login   string
	HtmlURL string `json:"html_url"`
}

type Repository struct {
	Name      string
	CreatedAt time.Time `json:"created_at"`
	HtmlURL   string    `json:"html_url"`
	Owner     *Owner
}

const temp = `{{  .TotalCount }} repositories:
{{ range .Items }}---------------------------------
Name: {{ .Name }}
Owner: {{ .Owner.Login }}
Created_at: {{ .CreatedAt | daysAgo }} days
{{ end }}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
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

func parseHTML() {
	var report = template.Must(template.New("repositories").Parse(`
<h1>{{ .TotalCount }} repositories </h1>
<table>
<tr style='text-align: left'>
<th>Name</th>
<th>Owner</th>
<th>Created_at</th>
</tr>
{{range .Items}}
<tr>
<td>{{ .Name }}</td>
<td>{{ .Owner.Login }}</td>
<td>{{ .CreatedAt }} </td>
</tr>
{{ end }}
</table>
`))
	result, err := SearchRepositories([]string{"vue"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func parseTemplate() {
	var report = template.Must(template.New("repositories").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp))

	result, err := SearchRepositories([]string{"vue"})
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
func main() {
	parseTemplate()
	parseHTML()
}
