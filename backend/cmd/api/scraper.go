package main

import (
	"context"
	"database/sql"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/internal/parser"
	"log"
	"strings"
	"sync"
	"time"
)

const fetchInterval = 3 * time.Hour
const fetchLimit = 2

func (app *application) fetchFeed(url string) (parser.Feed, error) {
	var feed parser.Feed

	resp, err := app.httpClient.Get(url)
	if err != nil {
		return feed, err
	}
	defer resp.Body.Close()

	feed, err = parser.NewParser().Parse(resp.Body)
	if err != nil {
		return feed, err
	}
	return feed, nil
}

func (app *application) ContinuousFeedScraping() {
	ticker := time.NewTicker(fetchInterval)
	for range ticker.C {
		feedsToFetch, err := app.db.GetNextFeedsToFetch(context.Background(), fetchLimit)
		if err != nil {
			log.Printf("Feed Fetcher could not get the next feeds to fetch: %v\n", err)
			continue
		}

		wg := sync.WaitGroup{}
		wg.Add(len(feedsToFetch))
		for _, feed := range feedsToFetch {
			go func() {
				defer wg.Done()

				err := app.db.MarkFeedFetched(context.Background(), feed.ID)
				if err != nil {
					log.Printf("Feed Fetcher could not mark the feed as fetched: %v\n", err)
					return
				}

				fetchedFeed, err := app.fetchFeed(feed.Url)
				if err != nil {
					log.Printf("Feed Fetcher could not get the feed %v: %v\n", feed.Url, err)
					return
				}

				log.Printf("Fetched feed: %v\n", feed.Url)
				var newPosts []database.Post
				for _, feedItem := range fetchedFeed.Items {
					post, err := app.db.CreatePost(context.Background(), database.CreatePostParams{
						Title: feedItem.Title,
						Url:   feedItem.Link,
						Description: sql.NullString{
							String: feedItem.Description,
							Valid:  true,
						},
						PublishedAt: sql.NullTime{
							Time:  feedItem.PublishedAt,
							Valid: true,
						},
						FeedID: feed.ID,
					})
					if err != nil {
						log.Printf("Feed Fetcher could not create the post: %v\n", err)
					} else {
						newPosts = append(newPosts, post)
					}
				}

				if err = app.applyFeedFilters(feed.ID, newPosts); err != nil {
					log.Printf("Feed Fetcher could not apply the feed filters: %v\n", err)
				}
			}()
		}
		wg.Wait()
	}
}

func (app *application) applyFeedFilters(feedID int64, newPosts []database.Post) error {
	feedFilters, err := app.db.GetFeedFilters(context.Background(), feedID)
	if err != nil {
		return err
	}

	for _, feedFilter := range feedFilters {
		filterText := strings.ToLower(feedFilter.FilterText)
		for _, post := range newPosts {
			title := strings.ToLower(post.Title)
			if strings.Contains(title, filterText) {
				_, err = app.db.CreatePostRead(context.Background(), database.CreatePostReadParams{
					UserID: feedFilter.UserID,
					PostID: post.ID,
				})
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
