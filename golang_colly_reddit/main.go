package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	StoryURL  string
	Source    string
	CrawledAt time.Time
	Comments  string
	Title     string
}

func main() {
	fmt.Println("Starting the application...")

	// Check if there are enough command-line arguments
	reddits := os.Args[1:]
	if len(reddits) == 0 {
		fmt.Println("Please provide at least one subreddit to crawl.")
		os.Exit(1)
	}

	stories := []item{}

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: old.reddit.com
		colly.AllowedDomains("old.reddit.com"),
		colly.Async(true),
	)

	// On every a element which has .top-matter attribute call callback
	// This class is unique to the div that holds all information about a story
	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		temp := item{}
		temp.StoryURL = e.ChildAttr("a[data-event-action=title]", "href")
		temp.Source = "https://old.reddit.com/r/programming/"
		temp.Title = e.ChildText("a[data-event-action=title]")
		temp.Comments = e.ChildAttr("a[data-event-action=comments]", "href")
		temp.CrawledAt = time.Now()
		fmt.Println(temp)
		stories = append(stories, temp)
	})

	// On every span tag with the class next-button
	c.OnHTML("span.next-button", func(h *colly.HTMLElement) {
		t := h.ChildAttr("a", "href")
		c.Visit(t)
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Error handling for HTTP requests
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Crawl all reddits the user passes in
	for _, reddit := range reddits {
		fmt.Println("Crawling", reddit)
		c.Visit(reddit)
	}
	fmt.Println("Crawling complete")

	c.Wait()
	fmt.Println(stories)
}
