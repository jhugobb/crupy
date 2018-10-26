package main

import (
	"log"
	"strconv"
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
					name := processCreate(msg[1:], &chars)
					response.Text = "The adventure of " + name + " has just begun!"
				} else {
					response.Text = err.Error()
				}
			case "showall":
				msg := strings.Split(update.Message.Text, " ")
				err := validateShowAll(msg[1:])
				if err == nil {
					response.Text = processShowAll(msg[1:], &chars)
				}
			default:
				response.Text = "Command not supported"
			}
			bot.Send(response)
		}
	}
}

func processCreate(msg []string, chars *Characters) string {
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

func processShowAll(msg []string, chars *Characters) string {
	s := strings.Builder{}
	switch msg[0] {
	case "dnd":
		dndcs := chars.dndcs
		for index, value := range dndcs {
			log.Print("I got in at least once")
			s.WriteString("Character #" + strconv.Itoa(index) + ": " + value.name + " is a [race] [class]\n")
		}
	default:
		return "neverhere"
	}
	return s.String()
}
