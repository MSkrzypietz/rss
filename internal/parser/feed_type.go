package parser

import (
	"bytes"
	"encoding/xml"
)

type feedType int

const (
	feedTypeUnknown feedType = iota
	feedTypeRss
	feedTypeAtom
)

func identifyFeedType(b []byte) (feedType, error) {
	buf := bytes.NewBuffer(b)
	decoder := xml.NewDecoder(buf)
	for {
		tok, err := decoder.Token()
		if err != nil {
			return feedTypeUnknown, err
		}

		switch elem := tok.(type) {
		case xml.StartElement:
			if elem.Name.Local == "rss" {
				return feedTypeRss, nil
			} else if elem.Name.Local == "feed" {
				return feedTypeAtom, nil
			}
		}
	}
}
