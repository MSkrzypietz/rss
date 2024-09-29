package parser

import (
	"bytes"
	"encoding/xml"
	"log"
	"time"
)

type rssFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Items []struct {
			Title       string  `xml:"title"`
			Link        string  `xml:"link"`
			Description string  `xml:"description"`
			PublishedAt *string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func parseRssFeed(b []byte) (Feed, error) {
	var feed rssFeed
	buf := bytes.NewBuffer(b)
	decoder := xml.NewDecoder(buf)
	err := decoder.Decode(&feed)
	if err != nil {
		return Feed{}, err
	}

	var items []Item
	for _, item := range feed.Channel.Items {
		publishedAt, err := parseRssPublishDate(item.PublishedAt)
		if err != nil {
			log.Printf("Rss parser could not parse the published date %v: %v\n", item.PublishedAt, err)
			continue
		}

		items = append(items, Item{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
		})
	}

	return Feed{
		Title: feed.Channel.Title,
		Items: items,
	}, nil
}

func parseRssPublishDate(date *string) (time.Time, error) {
	if date == nil {
		return time.Now(), nil
	}
	return time.Parse(time.RFC1123, *date)
}
