package telegrambotapi

import (
	"MIET-TelegramBot/internal/app/filesapi"
	"MIET-TelegramBot/internal/app/store/repository"
	"MIET-TelegramBot/internal/app/telegramserver"
	"MIET-TelegramBot/internal/app/timeapi"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	BotAPI         *tgbotapi.BotAPI
	UpdatesChannel tgbotapi.UpdatesChannel
	UniversityData *filesapi.ScheduleUniversity
	TimeInfo       *timeapi.TimeInformation
	DataBase       *repository.Repository
}

func CreateTelegramBot(configData *telegramserver.Config, resourcesPath string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(configData.Token)
	if err != nil {
		log.Println("Error get bot ", err.Error())
		return nil, err
	}

	universityData, err := filesapi.CreateScheduleUniversity(resourcesPath)

	if err != nil {
		log.Println("Error read resources files ", err.Error())
		return nil, err
	}

	timeInfo := timeapi.CreateTimeInformation()

	db, err := telegramserver.NewDB(configData)

	if err != nil {
		log.Println("Error create data base ", err.Error())
		return nil, err
	}

	repository := repository.NewRepository(db)

	return &TelegramBot{
		BotAPI:         bot,
		UniversityData: universityData,
		TimeInfo:       timeInfo,
		DataBase:       repository,
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
		//var responceToUser string
		if update.Message != nil { // If we got a message
			fmt.Println("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				fmt.Println("This is command : ", update.Message.Command(), update.Message.From.UserName)
				go b.handlersCommands(update.Message)
				continue
			}

			fmt.Println("This is message : ", update.Message.Text, update.Message.From.UserName, update.Message.From.ID)
		}
	}
}
