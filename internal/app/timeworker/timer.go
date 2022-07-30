package timeworker

import (
	"MIET-TelegramBot/internal/app/datareader"
	"fmt"
	"log"
	"time"
)

const (
	FirstNumerator    = "1 числитель"
	SecondNumerator   = "2 числитель"
	FirstDeniminator  = "1 знаменатель"
	SecondDeniminator = "2 знаменатель"
)

type WeekInformation struct {
	StartData  time.Time
	WeeksTypes map[int]string
}

type Timer struct {
	WeekInfo *WeekInformation
}

// func createWeekInformation() *WeekInformation {
// 	weekTypes := make(map[int]string)
// 	weekTypes[1] = FirstNumerator
// }

// Sunday Weekday = iota // Воскресенье
// Monday                // Понедельник
// Tuesday               // Вторник
// Wednesday             // Среда
// Thursday              // Четверг
// Friday                // Пятница
// Saturday              // Суббота

func (timer *Timer) GetTodayDayNumber() (string, int) {
	data := time.Now().Weekday()
	return data.String(), int(data)
}

func (timer *Timer) GetTomorrowDayNumber() (string, int) {
	data := time.Now().Add(24 * time.Hour).Weekday()
	return data.String(), int(data)
}

func (timer *Timer) GetCurrentWeek() {

}

func IdentifyCurrentPair(timeClass map[int8]datareader.TimeClasses) (string, error) {
	currentTime := time.Date(1, 1, 1, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)

	for i := 0; i < len(timeClass); i++ {
		timeFrom, err := convertStringToTime(timeClass[int8(i)].TimeFrom)
		if err != nil {
			log.Println(fmt.Sprint("IdentifyCurrentPair Error : convert time from string to date format failed %s", err.Error()))
			return "", err
		}

		timeTo, err := convertStringToTime(timeClass[int8(i)].TimeTo)

		if err != nil {
			log.Println(fmt.Sprint("IdentifyCurrentPair Error : convert time from string to date format failed %s", err.Error()))
			return "", err
		}

		if currentTime.Before(timeFrom) {
			log.Println("Before : ", timeFrom)
			hours, minutes, seconds := timeFrom.Clock()
			return fmt.Sprintf("Следующая пара %d в %d:%d:%d", i+1, hours, minutes, seconds), nil
		} else if currentTime.After(timeFrom) && currentTime.Before(timeTo) {
			hours, minutes, seconds := timeTo.Clock()
			return fmt.Sprintf("Сейчас идет пара %d до %d:%d:%d", i+1, hours, minutes, seconds), nil
		}
	}

	return "Сейчас пар нет", nil
}

func convertStringToTime(timeStr string) (time.Time, error) {
	timeT, err := time.Parse(time.RFC3339, timeStr+"Z")

	if err != nil {
		log.Println(fmt.Sprintf("Error convert string time to time type %s %s", err.Error(), timeStr))
		return timeT, err
	}

	return timeT, nil
}
