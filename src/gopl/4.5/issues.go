package main

import (
	"fmt"
	"gopl/4.5/github"
	"log"
	"os"
	"time"
)

const thirtyDays = 30 * 24 * 60 * 60
const oneYears = 365 * 24 * 60 * 60

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	return
	fmt.Printf("%d issues:\n", result.TotalCount)
	var beforeOy []*github.Issue
	var beforeOm []*github.Issue
	var beforeOn []*github.Issue
	now := time.Now()
	om := time.Unix(now.Unix()-thirtyDays, 0)
	oy := time.Unix(now.Unix()-oneYears, 0)

	for _, item := range result.Items {
		if item.CreatedAt.Before(oy) {
			beforeOy = append(beforeOy, item)
		}
		if item.CreatedAt.Before(om) && item.CreatedAt.After(oy) {
			beforeOm = append(beforeOm, item)
		}
		if item.CreatedAt.After(om) {
			beforeOn = append(beforeOn, item)
		}
	}
	fmt.Printf("一年前有 	%d issues:\n", len(beforeOy))
	for _, item := range beforeOy {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("不到一年有 	%d issues:\n", len(beforeOm))
	for _, item := range beforeOm {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
	fmt.Printf("不到一个月有  %d issues:\n", len(beforeOn))
	for _, item := range beforeOn {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
