package services

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"shopinbiz.telegram.bot/keyboards"
	"shopinbiz.telegram.bot/repositories"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Hi, wellcome from shopinbiz support bot"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func SetTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Please, write issue"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "ISSUE successfully created"
	err := repositories.SetTask(update)
	if err != nil {
		text = "Could not create issue"
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func DeleteTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	data, _ := repositories.GetAllTask(int16(update.Message.Chat.ID))
	var btns []tgbotapi.InlineKeyboardButton
	for i := 0; i < len(data); i++ {
		btn := tgbotapi.NewInlineKeyboardButtonData(data[i].Task, "delete_task="+data[i].ID.String())
		btns = append(btns, btn)
	}
	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i < len(btns) && i+1 < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
			rows = append(rows, row)
		} else if i < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i])
			rows = append(rows, row)
		}
	}
	fmt.Println(len(rows))
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
	text := "Please, select issue"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func DeleteTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, taskId string) {
	text := "Issue successfully deleted"
	err := repositories.DeleteTask(taskId)
	if err != nil {
		text = "Could not delete issue"
	}
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func ShowAllIssues(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Issues: \n"
	tasks, err := repositories.GetAllTask(int16(update.Message.Chat.ID))
	if err != nil {
		text = "Could not get issues"
	}
	for i := 0; i < len(tasks); i++ {
		text += tasks[i].Task + "\n"
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
