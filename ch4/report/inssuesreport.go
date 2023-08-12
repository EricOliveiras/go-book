package main

import (
	"gobook/ch4/github"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var formatIssues github.IssuesSearchResult
	var count int
	for _, item := range result.Items {
		if daysAgo(item.CreatedAt) <= 100 {
			formatIssues.Items = append(formatIssues.Items, item)
			count++
		}
	}
	formatIssues.TotalCount = count

	if err := report.Execute(os.Stdout, formatIssues); err != nil {
		log.Fatal(err)
	}
}
