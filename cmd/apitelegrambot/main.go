package main

import (
	"MIET-TelegramBot/internal/app/datareader"
	telegrambotapi "MIET-TelegramBot/internal/app/handlers"
	"MIET-TelegramBot/internal/app/telegramserver"
	"flag"
	"fmt"
	"log"
	"os"

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

	scheduleUniversity, err := datareader.CreateScheduleUniversity(config.ResourcesPath)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Result : ", scheduleUniversity.GroupsSchedule["ПИН-44"])

	bot, err := telegrambotapi.CreateTelegramBot(config.Token)

	if err != nil {
		fmt.Println("Create telegram bot failed")
		return
	}

	bot.ConfigTelegramBot()
	bot.StartTelegramBotServer()
}
