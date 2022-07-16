package main

import (
	telegrambotapi "MIET-TelegramBot/internal/app/telegram"
	"MIET-TelegramBot/internal/app/telegramserver"
	"flag"
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
		log.Fatal(err)
	}

	// fmt.Println("Result : ", scheduleUniversity.GroupsSchedule["ПИН-44"])
	// fmt.Println("Schedule classes : ", scheduleUniversity.GetClassTime())

	bot, err := telegrambotapi.CreateTelegramBot(config.Token, config.ResourcesPath)

	if err != nil {
		log.Fatal("Create telegram bot failed")
		return
	}

	bot.ConfigTelegramBot()
	bot.StartTelegramBotServer()
}
