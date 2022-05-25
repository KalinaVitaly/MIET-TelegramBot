package main

import (
	"MIET-TelegramBot/internal/app/recipientdata"
	"MIET-TelegramBot/internal/app/telegramserver"
	"flag"
	"github.com/BurntSushi/toml"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	configsPath string
)

func init() {
	flag.StringVar(&configsPath, "config-path", "configs/telegrambot.toml", "Path to configs")
}

func main() {
	flag.Parse()
	config := telegramserver.NewConfig()
	if _, err := toml.DecodeFile(configsPath, config); err != nil {
		log.Panic(err)
	}

	recipientdata.MakeRequest()

	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			var msg tgbotapi.MessageConfig
			log.Printf("[%s] %s %s", update.Message.Command(), update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Its a command")
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			}

			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
