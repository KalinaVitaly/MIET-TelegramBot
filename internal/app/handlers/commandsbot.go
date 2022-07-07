package telegrambotapi

import "MIET-TelegramBot/internal/app/datareader"

type CommandBot struct {
	Command  string
	UserName string
}

func CreateNewCommand(command string, userName string) *CommandBot {
	return &CommandBot{
		Command:  command,
		UserName: userName,
	}
}

func (cb *CommandBot) commandIdentification() string {
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
	switch cb.Command {

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
	case "help":
		return datareader.GetHelpMessage()
	case "class_time":
		responceClassTime := datareader.GetInstanceScheduleUniversity()
		return responceClassTime.ClassTimeToString()
	}

	return ""
}
