package telegrambotapi

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	BotAPI         *tgbotapi.BotAPI
	UpdateConfigTB tgbotapi.UpdateConfig
}

func CreateTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{BotAPI: bot}, nil
}

func (b *TelegramBot) ConfigTelegramBot() {
	b.BotAPI.Debug = true

	fmt.Println("Authorized on account %s", b.BotAPI.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	b.UpdateConfigTB = u
}

func (b *TelegramBot) StartTelegramBotServer() {
	updates := b.BotAPI.GetUpdatesChan(b.UpdateConfigTB)

	for update := range updates {
		if update.Message != nil { // If we got a message
			fmt.Println("[%s] %s", update.Message.From.UserName, update.Message.Text)
			fmt.Println("My output : ", update.Message.IsCommand())
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			b.BotAPI.Send(msg)
		}
	}
}
