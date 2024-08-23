package clients

import (
	"log"

	"shopinbiz.telegram.bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.Config("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	return bot
}
