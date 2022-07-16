package telegrambotapi

import (
	"MIET-TelegramBot/internal/app/datareader"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	BotAPI         *tgbotapi.BotAPI
	UpdatesChannel tgbotapi.UpdatesChannel
	UniversityData *datareader.ScheduleUniversity
}

func CreateTelegramBot(token string, resourcesPath string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Error get bot ", err.Error())
		return nil, err
	}

	universityData, err := datareader.CreateScheduleUniversity(resourcesPath)

	if err != nil {
		log.Println("Error read resources files ", err.Error())
		return nil, err
	}

	return &TelegramBot{
		BotAPI:         bot,
		UniversityData: universityData,
	}, nil
}

func (b *TelegramBot) ConfigTelegramBot() {
	b.BotAPI.Debug = true

	fmt.Println("Authorized on account %s", b.BotAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	b.UpdatesChannel = b.BotAPI.GetUpdatesChan(u)
}

func (b *TelegramBot) StartTelegramBotServer() {

	for update := range b.UpdatesChannel {
		var responceToUser string
		if update.Message != nil { // If we got a message
			fmt.Println("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				fmt.Println("This is command : ", update.Message.Command(), update.Message.From.UserName)
				b.handlersCommands(update.Message)
				continue
			}

			fmt.Println("This is message : ", update.Message.Text, update.Message.From.UserName)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, responceToUser)
			msg.ReplyToMessageID = update.Message.MessageID

			b.BotAPI.Send(msg)
		}
	}
}
