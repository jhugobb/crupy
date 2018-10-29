package main

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func processMessage(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range updates {
		if update.Message.IsCommand() {
			response := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "create":
				msg := strings.Split(update.Message.Text, " ")
				err := validateCreate(msg[1:])
				if err == nil {
					name := processCreate(msg[1:])
					response.Text = "The adventure of " + name + " has just begun!"
				} else {
					response.Text = err.Error()
				}
			case "showall":
				msg := strings.Split(update.Message.Text, " ")
				err := validateShowAll(msg[1:])
				if err == nil {
					response.Text = processShowAll(msg[1:])
					if len(response.Text) == 0 {
						response.Text = "I don't have any characters saved"
					}
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

func processCreate(msg []string) string {
	switch msg[0] {
	case "dnd":
		c := processDndCreate(msg[1])
		saveDndCharacter(c)
		return c.name
	}
	return "neverhere"
}

func processDndCreate(name string) DnDCharacter {
	return DnDCharacter{name: name}
}

func processShowAll(msg []string) string {
	var result string
	switch msg[0] {
	case "dnd":
		result = buildAllDnd()
	default:
		return "neverhere"
	}
	return result
}
