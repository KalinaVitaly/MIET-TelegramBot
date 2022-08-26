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
	//add error handle
	case "now":
		b.middlewareHandler(message, b.handleNowCommand)
	case "today":
		b.middlewareHandler(message, b.handleTodayCommand)
	case "tomorrow":
		b.middlewareHandler(message, b.handleTomorrowCommand)
	case "teacher_all":
		b.middlewareHandler(message, b.handleTeacherAllCommand)
	case "weekschedule":
		b.middlewareHandler(message, b.handleWeekScheduleCommand)
	case "weekschedule_short":
		b.middlewareHandler(message, b.handleWeekScheduleShortCommand)
	case "week":
		b.middlewareHandler(message, b.handleWeekCommand)
	case "group":
		b.middlewareHandler(message, b.handleGroupCommand)
	case "auth":
		b.middlewareHandler(message, b.handleAuthCommand)
	case "auth_teacher":
		b.middlewareHandler(message, b.handleAuthTeacherCommand)
	case "deauth":
		b.middlewareHandler(message, b.handleDeauthCommand)
	case "deauth_teacher":
		b.middlewareHandler(message, b.handleDeauthTeacherCommand)
	case "subscription":
		b.middlewareHandler(message, b.handleSubscribtionCommand)
	case "subscribe":
		b.middlewareHandler(message, b.handleSubscribeCommand)
	case "desubscibe":
		b.middlewareHandler(message, b.handleDesubscribeCommand)
	case "help":
		b.middlewareHandler(message, b.handleHelpCommand)
	case "class_time":
		b.middlewareHandler(message, b.handleClassTimeCommand)
	default:
		return b.handleDefaultCommand(message)
	}

	return nil
}

func (b *TelegramBot) handleNowCommand(message *tgbotapi.Message) error {
	msgText, err := b.TimeInfo.IdentifyCurrentPair(b.UniversityData.ClassTime)
	if err != nil {
		log.Println(fmt.Sprintf("Error now command : %s", err.Error()))
		return nil
	}
	log.Println(fmt.Sprintf("Get data now %s", msgText))
	return b.sendResponseMsg(message, msgText)
}

func (b *TelegramBot) handleTodayCommand(message *tgbotapi.Message) error {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")

	group, err := b.DataBase.User().Group(user)

	if err != nil {
		log.Println("Error get user group")
		return b.sendResponseMsg(message, "Что-то пошло не так при загрузке данных...")
	}

	weekType := b.TimeInfo.GetCurrentWeekType().WeekTypeNumber
	_, dayNumber := b.TimeInfo.GetTodayDayNumber()

	_, groupRus, isValidGroup := user.ValidGroup(group)

	if !isValidGroup {
		log.Println("Invalid group name")
		return b.sendResponseMsg(message, "Что-то не так с группой...")
	}

	_todaySchedule, err := b.UniversityData.GetClassesInSelectedDay(groupRus, dayNumber, weekType)

	fmt.Println("Day number ", dayNumber)
	if err != nil {
		log.Panicln(fmt.Sprintf("Error get classes in selected day %s", err.Error()))
		return b.sendResponseMsg(message, "Что-то пошло не так...")
	}

	if _todaySchedule == "" {
		_todaySchedule = "Сегодня пар нет"
	}

	return b.sendResponseMsg(message, _todaySchedule)
}

func (b *TelegramBot) handleTomorrowCommand(message *tgbotapi.Message) error {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")

	group, err := b.DataBase.User().Group(user)

	if err != nil {
		log.Println("Error get user group")
		return b.sendResponseMsg(message, "Что-то пошло не так при загрузке данных...")
	}

	_, groupRus, isValidGroup := user.ValidGroup(group)

	if !isValidGroup {
		log.Println("Invalid group name")
		return b.sendResponseMsg(message, "Что-то не так с группой...")
	}

	_, dayNumber, _, weekType := b.TimeInfo.GetTomorrowDayNumberAndWeekType()

	_todaySchedule, err := b.UniversityData.GetClassesInSelectedDay(groupRus, dayNumber, weekType)

	if err != nil {
		log.Panicln(fmt.Sprintf("Error get classes in selected day %s", err.Error()))
		return b.sendResponseMsg(message, "Что-то пошло не так...")
	}

	if _todaySchedule == "" {
		_todaySchedule = "Завтра пар нет"
	}

	return b.sendResponseMsg(message, _todaySchedule)
}

func (b *TelegramBot) handleTeacherAllCommand(message *tgbotapi.Message) error {

	return nil
}

func (b *TelegramBot) handleWeekScheduleCommand(message *tgbotapi.Message) error {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")

	group, err := b.DataBase.User().Group(user)

	if err != nil {
		log.Println("Error get user group")
		return b.sendResponseMsg(message, "Что-то пошло не так при загрузке данных...")
	}

	_, groupRus, isValidGroup := user.ValidGroup(group)

	if !isValidGroup {
		log.Println("Invalid group name")
		return b.sendResponseMsg(message, "Что-то не так с группой...")
	}

	weekInfo := b.TimeInfo.GetCurrentWeekType()
	weekSchedule := b.UniversityData.GetClassesCurrentWeek(groupRus, weekInfo.WeekTypeNumber)

	return b.sendResponseMsg(message, weekSchedule)
}

func (b *TelegramBot) handleWeekScheduleShortCommand(message *tgbotapi.Message) error {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")

	group, err := b.DataBase.User().Group(user)

	if err != nil {
		log.Println("Error get user group")
		return b.sendResponseMsg(message, "Что-то пошло не так при загрузке данных...")
	}

	_, groupRus, isValidGroup := user.ValidGroup(group)

	if !isValidGroup {
		log.Println("Invalid group name")
		return b.sendResponseMsg(message, "Что-то пошло не так при загрузке данных...")
	}

	weekInfo := b.TimeInfo.GetCurrentWeekType()
	weekSchedule, err := b.UniversityData.GetShortClassesCurrentWeek(groupRus, weekInfo.WeekTypeNumber)

	if err != nil {
		log.Println("Error get short current week classes")
		return b.sendResponseMsg(message, "Что-то пошло не так при загрузке данных...")
	}

	return b.sendResponseMsg(message, weekSchedule)
}

func (b *TelegramBot) handleWeekCommand(message *tgbotapi.Message) error {
	msgText := b.TimeInfo.GetCurrentWeekType()
	return b.sendResponseMsg(message, msgText.WeekTypeStr)
}

func (b *TelegramBot) handleGroupCommand(message *tgbotapi.Message) error {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")

	group, err := b.DataBase.User().Group(user)

	if err != nil {
		log.Println(fmt.Sprintf("Error get user group %+v from db", user))
		return err
	}

	return b.sendResponseMsg(message, group)
}

func (b *TelegramBot) handleAuthCommand(message *tgbotapi.Message) error {

	//Add group validation

	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")

	_, _, isGroupValid := user.ValidGroup(message.CommandArguments())
	if !isGroupValid {
		log.Println(fmt.Sprintf("Invalid group value : %s", message.CommandArguments()))
		return b.sendResponseMsg(message, "Группы не существует")
	}

	user.SetGroup(message.CommandArguments())

	if err := b.DataBase.User().Create(user); err != nil {
		log.Println(fmt.Sprintf("Error create user : %s", err.Error()))
		b.sendResponseMsg(message, "Ошибка при авторизации")
		return err
	}

	log.Println(fmt.Sprintf("New user auth : %+v", *user))
	return b.sendResponseMsg(message, "Пользователь успешно авторизован")
}

func (b *TelegramBot) handleAuthTeacherCommand(message *tgbotapi.Message) error {
	return nil
}

func (b *TelegramBot) handleDeauthCommand(message *tgbotapi.Message) error {
	user := models.CreateUserModel(message.From.ID, message.From.FirstName, message.From.LastName, message.From.UserName, "")

	if err := b.DataBase.User().Delete(user); err != nil {
		log.Println(fmt.Sprintf("Error create user : %s", err.Error()))
		b.sendResponseMsg(message, "Что-то пошло не так ...")
		return err
	}

	log.Println(fmt.Sprintf("User deauth : %+v", *user))
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
