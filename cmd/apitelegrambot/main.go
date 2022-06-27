package main

import (
	telegrambotapi "MIET-TelegramBot/internal/app/handlers"
	"MIET-TelegramBot/internal/app/telegramserver"
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
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

	bot, err := telegrambotapi.CreateTelegramBot(config.Token)

	if err != nil {
		fmt.Println("Create telegram bot failed")
		return
	}

	bot.ConfigTelegramBot()
	bot.StartTelegramBotServer()
}
