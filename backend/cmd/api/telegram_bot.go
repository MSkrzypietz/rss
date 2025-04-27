package main

import (
	"github.com/MSkrzypietz/rss/internal/database"
	"github.com/go-telegram/bot"
	"net/http"
)

func newBot(token string) (*bot.Bot, error) {
	return bot.New(token)
}

func (app *application) telegramEchoHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	var input struct {
		Message string `json:"message"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if !user.TelegramChatID.Valid {
		app.missingIntegrationResponse(w, r)
		return
	}

	_, err = app.telegramBot.SendMessage(r.Context(), &bot.SendMessageParams{
		ChatID: user.TelegramChatID.Int64,
		Text:   input.Message,
	})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	err = app.writeJSON(w, http.StatusOK, struct{}{}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
