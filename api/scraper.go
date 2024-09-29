package api

import (
	"context"
	"database/sql"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/internal/parser"
	"log"
	"sync"
	"time"
)

const fetchInterval = 3 * time.Hour
const fetchLimit = 2

func (cfg *Config) fetchFeed(url string) (parser.Feed, error) {
	var feed parser.Feed

	resp, err := cfg.httpClient.Get(url)
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

func (cfg *Config) ContinuousFeedScraping() {
	ticker := time.NewTicker(fetchInterval)
	for range ticker.C {
		feedsToFetch, err := cfg.db.GetNextFeedsToFetch(context.Background(), fetchLimit)
		if err != nil {
			log.Printf("Feed Fetcher could not get the next feeds to fetch: %v\n", err)
			continue
		}

		wg := sync.WaitGroup{}
		wg.Add(len(feedsToFetch))
		for _, feed := range feedsToFetch {
			go func() {
				defer wg.Done()

				err := cfg.db.MarkFeedFetched(context.Background(), feed.ID)
				if err != nil {
					log.Printf("Feed Fetcher could not mark the feed as fetched: %v\n", err)
					return
				}

				fetchedFeed, err := cfg.fetchFeed(feed.Url)
				if err != nil {
					log.Printf("Feed Fetcher could not get the feed %v: %v\n", feed.Url, err)
					return
				}

				log.Printf("Fetched feed: %v\n", feed.Url)
				for _, feedItem := range fetchedFeed.Items {
					_, err = cfg.db.CreatePost(context.Background(), database.CreatePostParams{
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
						log.Printf("Feed Fetcher could not get the post: %v\n", err)
					}
				}
			}()
		}
		wg.Wait()
	}
}
