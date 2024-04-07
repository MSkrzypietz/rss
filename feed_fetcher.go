package main

import "encoding/xml"

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

func (cfg *apiConfig) fetchFeeds(url string) ([]FetchedFeed, error) {
	var fetchedFeeds []FetchedFeed

	resp, err := cfg.httpClient.Get(url)
	if err != nil {
		return fetchedFeeds, err
	}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&fetchedFeeds)
	if err != nil {
		return fetchedFeeds, err
	}

	return fetchedFeeds, nil
}
