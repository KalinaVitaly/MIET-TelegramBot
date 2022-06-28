package main

import (
	"MIET-TelegramBot/internal/app/datareader"
	telegrambotapi "MIET-TelegramBot/internal/app/handlers"
	"MIET-TelegramBot/internal/app/telegramserver"
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configsPath   string
	resourcesPath string
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

	// scheduleData, err := datareader.GetFileNameFromDir(config.ResourcesPath)

	// if err != nil {
	// 	log.Panic(err)
	// }

	schedule, err := datareader.ReadDataFromFile("./resources/PIN-44.json")

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Data : ")

	for key, value := range schedule.Data {
		fmt.Println(key, value)
	}

	fmt.Println("Semestr : ")
	for key, value := range schedule.Semestr {
		fmt.Println(key, value)
	}

	fmt.Println("Times : ")
	for key, value := range schedule.Times {
		fmt.Println(key, value)
	}
	//fmt.Println("Data read :", schedule.Data[0].Group.Name)

	bot, err := telegrambotapi.CreateTelegramBot(config.Token)

	if err != nil {
		fmt.Println("Create telegram bot failed")
		return
	}

	bot.ConfigTelegramBot()
	bot.StartTelegramBotServer()
}
