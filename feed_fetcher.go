package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"log"
	"sync"
	"time"
)

const fetchInterval = 60 * time.Second
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

func (cfg *apiConfig) fetchFeed(url string) (FetchedFeed, error) {
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

func (cfg *apiConfig) continuousFeedFetcher() {
	ticker := time.NewTicker(fetchInterval)
	for range ticker.C {
		feedsToFetch, err := cfg.DB.GetNextFeedsToFetch(context.Background(), fetchLimit)
		if err != nil {
			log.Printf("Feed Fetcher could not get the next feeds to fetch: %v\n", err)
			continue
		}

		wg := sync.WaitGroup{}
		wg.Add(len(feedsToFetch))
		for _, feed := range feedsToFetch {
			go func() {
				defer wg.Done()

				err := cfg.DB.MarkFeedFetched(context.Background(), feed.ID)
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

					_, err = cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
						ID:    uuid.New(),
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
					var pqErr *pq.Error
					if err != nil && errors.As(err, &pqErr) && pqErr.Code.Name() != "unique_violation" {
						log.Printf("Feed Fetcher could not get the post: %v - %v\n", pqErr.Code.Name(), pqErr.Message)
					}
				}
			}()
		}
		wg.Wait()
	}
}
