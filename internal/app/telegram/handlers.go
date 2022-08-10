package telegrambotapi

import (
	"MIET-TelegramBot/internal/app/filesapi"
	"MIET-TelegramBot/internal/app/store/models"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *TelegramBot) handlersCommands(message *tgbotapi.Message) error {
	log.Println(fmt.Sprintf("Комманда %s от пользователя %s", message.Command(), message.From.UserName))
	switch message.Command() {

	case "now":
		return b.handleNowCommand(message)
	case "today":
		return b.handleTodayCommand(message)
	case "tomorrow":
		return b.handleTomorrowCommand(message)
	case "teacher_all":
		return b.handleTeacherAllCommand(message)
	case "weekschedule":
		return b.handleWeekScheduleCommand(message)
	case "weekschedule_short":
		return b.handleWeekScheduleShortCommand(message)
	case "week":
		return b.handleWeekCommand(message)
	case "group":
		return b.handleGroupCommand(message)
	case "auth":
		return b.handleAuthCommand(message)
	case "auth_teacher":
		return b.handleAuthTeacherCommand(message)
	case "deauth":
		return b.handleDeauthCommand(message)
	case "deauth_teacher":
		return b.handleDeauthTeacherCommand(message)
	case "subscription":
		return b.handleSubscribtionCommand(message)
	case "subscribe":
		return b.handleSubscribeCommand(message)
	case "desubscibe":
		return b.handleDesubscribeCommand(message)
	case "help":
		return b.handleHelpCommand(message)
	case "class_time":
		return b.handleClassTimeCommand(message)
	default:
		return b.handleDefaultCommand(message)
	}
}

func (b *TelegramBot) handleNowCommand(message *tgbotapi.Message) error {
	msgText, err := b.TimeInfo.IdentifyCurrentPair(b.UniversityData.ClassTime)
	if err != nil {
		log.Println(fmt.Sprint("Error now command : %s", err.Error()))
		return nil
	}
	log.Println(fmt.Sprintln("Get data now %s", msgText))
	return b.sendResponseMsg(message, msgText)
}

func (b *TelegramBot) handleTodayCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleTomorrowCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleTeacherAllCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleWeekScheduleCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleWeekScheduleShortCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleWeekCommand(message *tgbotapi.Message) error {
	msgText := b.TimeInfo.GetCurrentWeekType()
	return b.sendResponseMsg(message, msgText.WeekTypeStr)
}

func (b *TelegramBot) handleGroupCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleAuthCommand(message *tgbotapi.Message) error {

	//Add group validation
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, message.CommandArguments())

	if err := b.DataBase.User().Create(user); err != nil {
		log.Println(fmt.Sprint("Error create user : %s", err.Error()))
		b.sendResponseMsg(message, "Ошибка при авторизации")
		return err
	}

	log.Println(fmt.Sprint("New user auth : %+v", *user))
	return b.sendResponseMsg(message, "Пользователь успешно авторизован")
}

func (b *TelegramBot) handleAuthTeacherCommand(message *tgbotapi.Message) error {
	return nil
}

func (b *TelegramBot) handleDeauthCommand(message *tgbotapi.Message) error {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, message.CommandArguments())

	if err := b.DataBase.User().Delete(user); err != nil {
		log.Println(fmt.Sprint("Error create user : %s", err.Error()))
		b.sendResponseMsg(message, "Что-то пошло не так ...")
		return err
	}

	log.Println(fmt.Sprint("User deauth : %+v", *user))
	return b.sendResponseMsg(message, "Пользователь деавторизован")
}

func (b *TelegramBot) handleDeauthTeacherCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleSubscribtionCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleSubscribeCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleDesubscribeCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleHelpCommand(message *tgbotapi.Message) error {
	msgText := filesapi.GetHelpMessage()
	return b.sendResponseMsg(message, msgText)
}

func (b *TelegramBot) handleClassTimeCommand(message *tgbotapi.Message) error {
	msgText := b.UniversityData.ClassTimeToString()
	return b.sendResponseMsg(message, msgText)
}

func (b *TelegramBot) handleDefaultCommand(message *tgbotapi.Message) error {
	msgText := fmt.Sprintf("Неизвестная комманда %s", message.Command())
	return b.sendResponseMsg(message, msgText)
}

func (b *TelegramBot) sendResponseMsg(message *tgbotapi.Message, msgText string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, msgText)

	if _, err := b.BotAPI.Send(msg); err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	return nil
}
