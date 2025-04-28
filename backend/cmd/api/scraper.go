package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/MSkrzypietz/rss/internal/parser"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"net/http"
	"strings"
	"sync"
	"time"
)

const fetchInterval = 3 * time.Hour
const fetchLimit = 2

func (app *application) getFeedByUrl(url string) (parser.Feed, error) {
	var feed parser.Feed

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return feed, err
	}

	req.Header.Set("Accept", "application/rss+xml, application/xml;q=0.9, */*;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:136.0) Gecko/20100101 Firefox/136.0")

	resp, err := app.httpClient.Do(req)
	if err != nil {
		return feed, err
	}
	defer resp.Body.Close()

	feed, err = parser.NewParser(app.logger).Parse(resp.Body)
	if err != nil {
		return feed, err
	}
	return feed, nil
}

func (app *application) fetchFeed(feed database.Feed) {
	err := app.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		app.logger.Error("Feed Fetcher could not mark the feed as fetched", "error", err)
		return
	}

	fetchedFeed, err := app.getFeedByUrl(feed.Url)
	if err != nil {
		app.logger.Error("Feed Fetcher could not get the feed", "feed", feed.Url, "error", err)
		return
	}

	app.logger.Info("Fetched feed", "feed", feed.Url)
	var newPosts []database.Post
	for _, feedItem := range fetchedFeed.Items {
		post, err := app.db.CreatePost(context.Background(), database.CreatePostParams{
			Title: feedItem.Title,
			Url:   feedItem.Link,
			Description: sql.NullString{
				String: feedItem.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  feedItem.PublishedAt,
				Valid: true,
			},
			FeedID: feed.ID,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "UNIQUE constraint") {
				app.logger.Error("Feed Fetcher could not create the post", "error", err)
			}
		} else {
			newPosts = append(newPosts, post)
		}
	}

	filteredUserPosts := app.applyFeedFilters(feed.ID, newPosts)

	feedFollowers, err := app.db.GetFeedFollowers(context.Background(), feed.ID)
	if err != nil {
		app.logger.Error("Feed Fetcher could not get the followers for the feed", "error", err)
		return
	}

	for _, feedFollowerID := range feedFollowers {
		msg := ""
		for _, post := range newPosts {
			if _, ok := filteredUserPosts[feedFollowerID]; ok {
				if _, ok := filteredUserPosts[feedFollowerID][post.ID]; ok {
					continue
				}
			}
			msg += fmt.Sprintf("\n\\- [%s](%s)", bot.EscapeMarkdown(post.Title), post.Url)
		}

		if msg == "" {
			continue
		}
		msg = "New posts:" + msg

		user, err := app.db.GetUserByID(context.Background(), feedFollowerID)
		if err != nil {
			app.logger.Error("Feed Fetcher could not get the user", "user", feedFollowerID, "error", err)
			continue
		}

		_, err = app.telegramBot.SendMessage(context.Background(), &bot.SendMessageParams{
			ChatID:    user.TelegramChatID.Int64,
			Text:      msg,
			ParseMode: models.ParseModeMarkdown,
			LinkPreviewOptions: &models.LinkPreviewOptions{
				IsDisabled: pointTo(true),
			},
		})
		if err != nil {
			app.logger.Error("Feed Fetcher could not send a message to the user", "user", feedFollowerID, "error", err)
		}
	}
}

func (app *application) ContinuousFeedScraping() {
	ticker := time.NewTicker(fetchInterval)
	for range ticker.C {
		feedsToFetch, err := app.db.GetNextFeedsToFetch(context.Background(), fetchLimit)
		if err != nil {
			app.logger.Error("Feed Fetcher could not get the next feeds to fetch", "error", err)
			continue
		}

		wg := sync.WaitGroup{}
		wg.Add(len(feedsToFetch))
		for _, feed := range feedsToFetch {
			go func() {
				defer wg.Done()
				app.fetchFeed(feed)
			}()
		}
		wg.Wait()
	}
}

func (app *application) applyFeedFilters(feedID int64, newPosts []database.Post) map[int64]map[int64]struct{} {
	feedFilters, err := app.db.GetFeedFilters(context.Background(), feedID)
	if err != nil {
		return nil
	}

	filteredUserPosts := make(map[int64]map[int64]struct{})
	for _, feedFilter := range feedFilters {
		filterText := strings.ToLower(feedFilter.FilterText)
		for _, post := range newPosts {
			title := strings.ToLower(post.Title)
			if strings.Contains(title, filterText) {
				_, err = app.db.CreatePostRead(context.Background(), database.CreatePostReadParams{
					UserID: feedFilter.UserID,
					PostID: post.ID,
				})
				if err != nil {
					app.logger.Error("Feed Fetcher could not create the post", "error", err)
					continue
				}
				if _, ok := filteredUserPosts[feedFilter.UserID]; !ok {
					filteredUserPosts[feedFilter.UserID] = make(map[int64]struct{})
				}
				filteredUserPosts[feedFilter.UserID][post.ID] = struct{}{}
			}
		}
	}
	return filteredUserPosts
}
