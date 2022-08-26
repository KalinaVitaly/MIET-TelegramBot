package telegrambotapi

import (
	"MIET-TelegramBot/internal/app/store/models"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *TelegramBot) middlewareHandler(message *tgbotapi.Message, handler func(*tgbotapi.Message) error) {
	log.Println(fmt.Sprintf("Комманда %s от пользователя %s", message.Command(), message.From.UserName))

	//added statistic geter

	if isCommandNeedingAuth(message.Command()) {
		isAuth, msg, err := b.isUserAuth(message)
		if err != nil {
			b.sendResponseMsg(message, msg)
			return
		} else if !isAuth {
			b.sendResponseMsg(message, msg)
			return
		}
	}

	handler(message)
}

func (b *TelegramBot) isUserAuth(message *tgbotapi.Message) (bool, string, error) {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")
	if isAuth, err := b.DataBase.User().Contains(user); err != nil {
		log.Println(fmt.Sprintf("Error check contains user in db : %s", err.Error()))
		return false, "Что-то пошло не так...", err
	} else {
		if isAuth {
			log.Println(fmt.Sprintf("User %+v contains in db", user))
			return true, "Пользователь авторизован", nil
		}
	}

	return false, "Пользователь не авторизован", nil
}

func isCommandNeedingAuth(command string) bool {
	commands := [...]string{
		"today",
		"tomorrow",
		"weekschedule",
		"weekschedule_short",
		"group",
		"deauth",
		"deauth_teacher",
		"subscription",
		"subscribe",
		"desubscibe",
	}

	for _, value := range commands {
		if value == command {
			return true
		}
	}

	return false
}
