package filesapi

import (
	"encoding/json"
	"log"
)

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

type TimeClasses struct {
	Time     string `json:"Time"`
	TimeFrom string `json:"TimeFrom"`
	TimeTo   string `json:"TimeTo"`
}

type RoomClass struct {
	Name string `json:"Name"`
}

type GroupData struct {
	Name string `json:"Name"`
}

type ClassData struct {
	Name        string `json:"Name"`
	TeacherFull string `json:"TeacherFull"`
	Teacher     string `json:"Teacher"`
	Form        string `json:"Form"`
}

type GroupSchedule struct {
	Times []TimeClasses `json:"Times"`
	Data  []struct {
		Day       int         `json:"Day"`
		DayNumber int         `json:"DayNumber"`
		Time      TimeClasses `json:"Time"`
		Class     ClassData   `json:"Class"`
		Group     GroupData   `json:"Group"`
		Room      RoomClass   `json:"Room"`
	} `json:"Data"`
	Semestr string `json:"Semestr"`
}

type fileTransferData struct {
	readError      error
	groupsSchedule *GroupSchedule
	filePath       string
}

func newFileTransferData(_groupsSchedule *GroupSchedule, _filePath string, _readError error) *fileTransferData {
	return &fileTransferData{
		readError:      _readError,
		groupsSchedule: _groupsSchedule,
		filePath:       _filePath,
	}
}

func (f *fileTransferData) containsError() bool {
	if f.readError != nil {
		return true
	}

	return false
}

func NewGroupSchedule(data []byte) (*GroupSchedule, error) {
	var groupSchedule GroupSchedule

	if err := json.Unmarshal(data, &groupSchedule); err != nil {
		log.Println("Error: Unmarshal group schedule failed!")
		return nil, err
	}

	return &groupSchedule, nil
}

func GetHelpMessage() string {
	return helpMessage
}
