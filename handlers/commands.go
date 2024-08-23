package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"shopinbiz.telegram.bot/services"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "set_issue":
		services.SetTask(bot, update)
	case "delete_issue":
		services.DeleteTask(bot, update)
	case "show_all_issue":
		services.ShowAllIssues(bot, update)
	}
}
