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

	c.OnHTML("div.commentarea div.comment", func(e *colly.HTMLElement) {
		comment := redditComment{
			CommentURL: e.Request.URL.String(),
			Source:     "reddit",
			CrawledAt:  time.Now(),
			Comment:    e.ChildText("div.usertext-body may-blank-within md-container"),
		}
		comments = append(comments, comment)
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
