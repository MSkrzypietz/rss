package parser

import (
	"bytes"
	"encoding/xml"
	"html"
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

func (p *Parser) parseRssFeed(b []byte) (Feed, error) {
	var feed rssFeed
	buf := bytes.NewBuffer(b)
	decoder := xml.NewDecoder(buf)
	err := decoder.Decode(&feed)
	if err != nil {
		return Feed{}, err
	}

	var items []Item
	for _, item := range feed.Channel.Items {
		publishedAt, err := parsePublishDate(item.PublishedAt)
		if err != nil {
			p.logger.Error("Rss parser could not parse the published date", "publishedAt", item.PublishedAt, "error", err)
			continue
		}

		items = append(items, Item{
			Title:       html.UnescapeString(item.Title),
			Link:        item.Link,
			Description: html.UnescapeString(item.Description),
			PublishedAt: publishedAt,
		})
	}

	return Feed{
		Title: html.UnescapeString(feed.Channel.Title),
		Items: items,
	}, nil
}
