package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/MathewKostiuk/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	categories := sortCategories(result)
	fmt.Printf("%d issues:\n", result.TotalCount)
	for key, category := range categories {
		fmt.Printf("%s\n", key)
		for _, item := range category {
			fmt.Printf("%v #%-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}
}

func sortCategories(result *github.IssuesSearchResult) map[string][]*github.Issue {
	sort.Slice(result.Items, func(i, j int) bool { return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt) })
	times := make(map[string][]*github.Issue)
	t := time.Now()
	lastMonth := t.AddDate(0, -1, 0)
	lastYear := t.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.After(lastMonth) {
			times["Less than a month"] = append(times["Less than a month"], item)
			continue
		}
		if item.CreatedAt.After(lastYear) {
			times["Less than one year"] = append(times["Less than one year"], item)
			continue
		}
		if item.CreatedAt.Before(lastYear) {
			times["More than one year"] = append(times["More than one year"], item)
			continue
		}
	}
	return times
}
