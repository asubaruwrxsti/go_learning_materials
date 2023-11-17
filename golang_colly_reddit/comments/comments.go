package comments

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type redditComment struct {
	CommentURL string
	Source     string
	CrawledAt  time.Time
	Comment    string
}

// crawls the comments of a reddit post
func CrawlRedditComments(c *colly.Collector, redditURL string) []redditComment {
	var comments []redditComment

	c.OnHTML("div.commentarea div.usertext-body div.md", func(e *colly.HTMLElement) {
		fmt.Println("Found comment:", e.Text)
		comments = append(comments, redditComment{
			CommentURL: e.Request.URL.String(),
			Source:     "reddit",
			CrawledAt:  time.Now(),
			Comment:    e.Text,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(redditURL)
	return comments
}
