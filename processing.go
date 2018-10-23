package main

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func processMessage(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, chars Characters) {
	for update := range updates {
		if update.Message.IsCommand() {
			response := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "create":
				msg := strings.Split(update.Message.Text, " ")
				err := validateCreate(msg[1:])
				if err == nil {
					name := processCreate(msg[1:], chars)
					response.Text = "The adventure of " + name + " has just begun!"
				} else {
					response.Text = err.Error()
				}
			default:
				response.Text = "Command not supported"
			}
			bot.Send(response)
		}
	}
}

func processCreate(msg []string, chars Characters) string {
	switch msg[0] {
	case "dnd":
		c := processDndCreate(msg[1])
		chars.addDnd(c)
		return c.name
	}
	return "neverhere"
}

func processDndCreate(name string) DnDCharacter {
	return DnDCharacter{name: name}
}
