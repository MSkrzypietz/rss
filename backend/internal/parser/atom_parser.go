package parser

import (
	"bytes"
	"encoding/xml"
	"html"
)

type atomFeed struct {
	Title   string `xml:"title"`
	Entries []struct {
		Title string `xml:"title"`
		Link  struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
		Summary     string  `xml:"summary"`
		PublishedAt *string `xml:"published"`
	} `xml:"entry"`
}

func (p *Parser) parseAtomFeed(b []byte) (Feed, error) {
	var feed atomFeed
	buf := bytes.NewBuffer(b)
	decoder := xml.NewDecoder(buf)
	err := decoder.Decode(&feed)
	if err != nil {
		return Feed{}, err
	}

	var items []Item
	for _, entry := range feed.Entries {
		publishedAt, err := parsePublishDate(entry.PublishedAt)
		if err != nil {
			p.logger.Error("Atom parser could not parse the published date %v: %v\n", "publishedAt", entry.PublishedAt, "error", err)
			continue
		}

		items = append(items, Item{
			Title:       html.UnescapeString(entry.Title),
			Link:        entry.Link.Href,
			Description: html.UnescapeString(entry.Summary),
			PublishedAt: publishedAt,
		})
	}

	return Feed{
		Title: html.UnescapeString(feed.Title),
		Items: items,
	}, nil
}
