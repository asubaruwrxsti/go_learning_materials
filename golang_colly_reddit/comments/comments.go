package comments

import (
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
func CrawlRedditComments(c *colly.Collector, redditURL string, defaultLimitComment int) ([]redditComment, error) {
	var comments []redditComment
	var err_ error
	limit_comment := 0

	c.OnHTML(".entry .usertext .usertext-body", func(e *colly.HTMLElement) {
		if limit_comment >= defaultLimitComment {
			return
		}
		comments = append(comments, redditComment{
			CommentURL: e.Request.URL.String(),
			Source:     "reddit",
			CrawledAt:  time.Now(),
			Comment:    e.Text,
		})
		limit_comment++
	})

	c.OnError(func(r *colly.Response, err error) {
		err_ = err
	})

	c.Visit(redditURL)
	c.Wait()
	return comments, err_
}
