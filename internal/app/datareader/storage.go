package datareader

import (
	"encoding/json"
	"fmt"
)

type TimeClass struct {
	Time     string `json:"Time"`
	Code     int    `json:"Code"`
	TimeFrom string `json:"TimeFrom"`
	TimeTo   string `json:"TimeTo"`
}

type GroupSchedule struct {
	Times []TimeClass `json:"Times"`
	Data  []struct {
		Day       int `json:"Day"`
		DayNumber int `json:"DayNumber"`
		Time      struct {
			Time     string `json:"Time"`
			Code     int    `json:"Code"`
			TimeFrom string `json:"TimeFrom"`
			TimeTo   string `json:"TimeTo"`
		} `json:"Time"`
		Class struct {
			Code        string `json:"Code"`
			Name        string `json:"Name"`
			TeacherFull string `json:"TeacherFull"`
			Teacher     string `json:"Teacher"`
			Form        string `json:"Form"`
		} `json:"Class"`
		Group struct {
			Code string `json:"Code"`
			Name string `json:"Name"`
		} `json:"Group"`
		Room struct {
			Code int    `json:"Code"`
			Name string `json:"Name"`
		} `json:"Room"`
	} `json:"Data"`
	Semestr string `json:"Semestr"`
}

var helpMessage string = `Список команд:
						/now (Сейчас) - Пара сейчас
						/today (Сегодня) - Расписание на сегодня
						/tomorrow (Завтра) - Расписание на завтра
						<День недели> - Расписание на конкретный день в рамках 
						текущей недели
						<День недели> <Тип недели> - Расписание на конкретный 
						день указанием типа недели в формате: 1ч, 1з, 2ч, 2з
						/teacher_all - Поиск преподавателя
						/weekschedule - Расписание на неделю
						/weekschedule_short - Расписание на неделю в кратком виде
						/week - Текущая неделя
						/group - Показать группу
						/auth - Авторизация
						/auth_teacher - Авторизация для преподавателя
						/deauth - Деавторизация
						/deauth_teacher - Деавторизация для преподавателя
						/subscription - Статус подписки
						/subscribe - Подписаться на уведомления
						/desubscribe - Отписаться от уведомлений
						=======================
						Для получения информации о расписании без авторизации 
						запросы:
						Сейчас <Группа>
						Сегодня <Группа>
						Завтра <Группа>
						<День недели> <Группа>
						<День недели> <Тип недели> <Группа>
						Обратная связь: @vibrunum
						@mieteventsbot`

func NewGroupSchedule(data []byte) (*GroupSchedule, error) {
	var groupSchedule GroupSchedule

	if err := json.Unmarshal(data, &groupSchedule); err != nil {
		fmt.Println("Error: Unmarshal group schedule failed!")
		return nil, err
	}

	return &groupSchedule, nil
}

func GetHelpMessage() string {
	return helpMessage
}
