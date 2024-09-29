package parser

import (
	"fmt"
	"io"
	"time"
)

type Item struct {
	Title       string
	Link        string
	Description string
	PublishedAt time.Time
}

type Feed struct {
	Title       string
	Description string
	Items       []Item
}

type Parser struct {
}

func NewParser() Parser {
	return Parser{}
}

func (p Parser) Parse(r io.Reader) (Feed, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return Feed{}, err
	}

	feedType, err := identifyFeedType(data)
	if err != nil {
		return Feed{}, err
	}

	switch feedType {
	case feedTypeRss:
		return parseRssFeed(data)
	case feedTypeAtom:
		return parseAtomFeed(data)
	default:
		return Feed{}, fmt.Errorf("unknown feed type")
	}
}
