package main

import (
	"context"
	"encoding/xml"
	"log"
	"sync"
	"time"
)

const fetchInterval = 30 * time.Second
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
					log.Println(feedItem.Title)
				}
			}()
		}
		wg.Wait()
	}
}
