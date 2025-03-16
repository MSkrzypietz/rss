package parser

import (
	"fmt"
	"io"
	"log/slog"
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
	logger *slog.Logger
}

func NewParser(logger *slog.Logger) *Parser {
	return &Parser{
		logger: logger,
	}
}

func (p *Parser) Parse(r io.Reader) (Feed, error) {
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
		return p.parseRssFeed(data)
	case feedTypeAtom:
		return p.parseAtomFeed(data)
	default:
		return Feed{}, fmt.Errorf("unknown feed type")
	}
}
