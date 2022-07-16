package telegrambotapi

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *TelegramBot) handlersCommands(message *tgbotapi.Message) error {
	// 	now - Пара сейчас
	// today - Расписание на сегодня
	// tomorrow - Расписание на завтра
	// teacher_all - Поиск преподавателя
	// weekschedule - Расписание на неделю
	// weekschedule_short - Краткое расписание
	// week - Текущая неделя
	// group - Показать группу
	// auth - Авторизация
	// auth_teacher - Авторизация для преподавателя
	// deauth - Деавторизация
	// deauth_teacher - Деавторизация для преподавателя
	// subscription - Статус подписки
	// subscribe - Подписаться на уведомления
	// desubscibe - Отписаться от уведомлений
	// help - Помощь

	switch message.Command() {

	case "now":

	case "today":

	case "tomorrow":

	case "teacher_all":

	case "weekschedule":

	case "weekschedule_short":

	case "week":

	case "group":

	case "auth":

	case "auth_teacher":
	case "deauth":
	case "deauth_teacher":
	case "subscription":
	case "subscribe":
	case "desubscibe":
	// case "help":
	// 	return datareader.GetHelpMessage()
	// case "class_time":
	// 	responceClassTime := datareader.GetInstanceScheduleUniversity()
	// 	return responceClassTime.ClassTimeToString()
	// }
	default:
		log.Println(fmt.Sprintf("Неизвестная комманда %s", message.Command()))
		return b.handleDefaultCommand(message)
	}

	return nil
}

func (b *TelegramBot) handleNowCommand(message *tgbotapi.Message) error {

	return nil
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

	return nil
}

func (b *TelegramBot) handleGroupCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleAuthCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleAuthTeacherCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleDeauthCommand(message *tgbotapi.Message) error {

	return nil
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

	return nil
}

func (b *TelegramBot) handleClassTimeCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleDefaultCommand(message *tgbotapi.Message) error {
	if _, err := b.BotAPI.Send(fmt.Sprintf("Неизвестная комманда %s", message.Command())); err != nil {
		return err
	}

	return nil
}
