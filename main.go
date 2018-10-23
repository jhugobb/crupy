package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	characters := Characters{}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://crupy-bot.herokuapp.com/" + bot.Token))

	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	processMessage(updates, bot, characters)
}

func dndCharacter(name string) DnDCharacter {
	character := DnDCharacter{name: name}
	return character
}
