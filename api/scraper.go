package api

import (
	"context"
	"database/sql"
	"encoding/xml"
	"github.com/MSkrzypietz/rss/internal/database"
	"log"
	"sync"
	"time"
)

const fetchInterval = 60 * time.Minute
const fetchLimit = 2

type FetchedFeed struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Items       []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			PublishedAt string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func (cfg *Config) fetchFeed(url string) (FetchedFeed, error) {
	var fetchedFeed FetchedFeed

	resp, err := cfg.httpClient.Get(url)
	if err != nil {
		return fetchedFeed, err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&fetchedFeed)
	if err != nil {
		return fetchedFeed, err
	}

	return fetchedFeed, nil
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

				log.Printf("Fetched feed %v: %v\n", feed.Url, fetchedFeed.Channel.Title)
				for _, feedItem := range fetchedFeed.Channel.Items {
					publishedAt, err := time.Parse(time.RFC1123, feedItem.PublishedAt)
					if err != nil {
						log.Printf("Feed Fetcher could not parse the published date %v: %v\n", feedItem.PublishedAt, err)
						continue
					}

					_, err = cfg.db.CreatePost(context.Background(), database.CreatePostParams{
						Title: feedItem.Title,
						Url:   feedItem.Link,
						Description: sql.NullString{
							String: feedItem.Description,
							Valid:  true,
						},
						PublishedAt: sql.NullTime{
							Time:  publishedAt,
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
