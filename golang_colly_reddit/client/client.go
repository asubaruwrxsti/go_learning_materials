package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"golang_colly_reddit/comments"

	"github.com/gocolly/colly"
)

type item struct {
	StoryURL    string
	Source      string
	CrawledAt   time.Time
	CommentsUrl string
	Title       string
}

func (i item) toString() string {
	return fmt.Sprintf("Source: %s\nTitle: %s\nStoryURL: %s\nCommentsUrl: %s\nCrawledAt: %s\n\n", i.Source, i.Title, i.StoryURL, i.CommentsUrl, i.CrawledAt)
}

func main() {
	fmt.Println("Starting the application...")

	// Check if there are enough command-line arguments
	reddits := os.Args[1:]
	if len(reddits) == 0 {
		fmt.Println("Please provide at least one subreddit to crawl.")
		os.Exit(1)
	}

	// Limit the number of posts to crawl
	defaultLimitPost := 1
	defaultLimitComment := 1

	switch len(os.Args) {
	case 1:
		fmt.Println("Please provide at least one subreddit to crawl.")
		os.Exit(1)
	case 2:
		fmt.Println("Using default number of posts to crawl:", defaultLimitPost)
	case 3:
		fmt.Println("Using custom number of posts to crawl:", os.Args[2])
	default:
		fmt.Println("Too many arguments. Please provide only one subreddit to crawl.")
		os.Exit(1)
	}

	limit_post := defaultLimitPost
	if len(os.Args) > 2 {
		// Convert the custom limit_post to an integer
		var err error
		limit_post, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Please provide a valid number of posts to crawl.")
			os.Exit(1)
		}
	}
	fmt.Println("")

	// var stories []item
	stories := []item{}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com"),
		colly.Async(true),
	)

	var count_post int = 0
	// On every a element which has .top-matter attribute call callback
	// This class is unique to the div that holds all information about a story
	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		if count_post >= limit_post {
			return
		}
		temp := item{}
		temp.StoryURL = e.ChildAttr("a[data-event-action=title]", "href")
		temp.Source = reddits[0]
		temp.Title = e.ChildText("a[data-event-action=title]")
		temp.CommentsUrl = e.ChildAttr("a[data-event-action=comments]", "href")
		temp.CrawledAt = time.Now()
		stories = append(stories, temp)
		count_post++
	})

	// On every span tag with the class next-button
	c.OnHTML("span.next-button", func(h *colly.HTMLElement) {
		if count_post >= limit_post {
			return
		}
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
		// fmt.Println("Visiting", r.URL.String())
		// fmt.Println("")
	})

	// Error handling for HTTP requests
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Crawl all reddits the user passes in
	for _, reddit := range reddits {
		c.Visit(reddit)
	}

	c.Wait()
	for _, story := range stories {
		fmt.Println(story.toString())
		storyComments, err_ := comments.CrawlRedditComments(c, story.CommentsUrl, defaultLimitComment)
		if err_ != nil {
			fmt.Println("Error crawling comments:", err_)
		}
		fmt.Println("Comments:")
		fmt.Println(storyComments)
		fmt.Println("")
	}
	defer fmt.Println("Crawling complete")
}
