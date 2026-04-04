package main

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/Sadzeih/valcompbot/config"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var (
	lastSENsince    time.Time
	lastSENsinceFmt = "The last Sentinels thread was %d days ago. 0 days since the last Sentinels thread."

	sentinelsRegex = regexp.MustCompile(`(?i)\b((SEN(tinels)?)|(n(a|4)rrate|redux(x)?|kyu|cortezia|zyto|kaplan|gunter|johnqt))\b`)
)

func DaysSinceLastSentinelsPost(c *reddit.Client) error {
	lastSENpost, err := LookForLastSentinelsPost(c)
	if err != nil {
		return err
	}
	if lastSENpost != nil {
		lastSENsince = lastSENpost.Created.Time
	}

	newPosts, errs, closeChan := c.Stream.Posts(config.Get().RedditSubreddit, reddit.StreamDiscardInitial)
	defer closeChan()

	for {
		select {
		case err := <-errs:
			return err
		case post := <-newPosts:
			if sentinelsRegex.MatchString(post.Title) {
				str := fmt.Sprintf(lastSENsinceFmt, int(time.Since(lastSENsince).Hours()/24))
				if lastSENsince.IsZero() && lastSENpost == nil {
					str = "I don't know when the last Sentinels post was (but it can't have been that long ago...). 0 days since the last Sentinels post."
				}
				com, _, err := c.Comment.Submit(context.Background(), post.FullID, str)
				if err != nil {
					return err
				}

				_, err = c.Moderation.DistinguishAndSticky(context.Background(), com.FullID)
				if err != nil {
					return err
				}
				lastSENsince = time.Now()
			}
		}
	}
}

func LookForLastSentinelsPost(c *reddit.Client) (*reddit.Post, error) {
	// Check 100 latest posts
	posts, _, err := c.Subreddit.NewPosts(context.Background(), config.Get().RedditSubreddit, &reddit.ListOptions{
		Limit: 100,
	})
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		if sentinelsRegex.MatchString(post.Title) {
			return post, nil
		}
	}

	if len(posts) < 100 {
		return nil, nil
	}

	// Check 300 more in case a thread has not been found
	for i := 0; i < 3; i++ {
		posts, _, err = c.Subreddit.NewPosts(context.Background(), config.Get().RedditSubreddit, &reddit.ListOptions{
			Limit:  100,
			Before: posts[len(posts)-1].FullID,
		})
		if err != nil {
			return nil, err
		}
		for _, post := range posts {
			if sentinelsRegex.MatchString(post.Title) {
				return post, nil
			}
		}
	}

	return nil, nil
}
