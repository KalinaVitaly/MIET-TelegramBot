package timeapi

import (
	"MIET-TelegramBot/internal/app/filesapi"
	"fmt"
	"log"
	"sync"
	"time"
)

const (
	weekTypesCount = 4
	Monday         = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var WeekTypes [weekTypesCount]string

func init() {
	WeekTypes = [...]string{
		"1 числитель",
		"2 числитель",
		"1 знаменатель",
		"2 знаменатель",
	}
}

type WeekInformation struct {
	weekTypeStr    string
	weekTypeNumber int
	mutex          sync.RWMutex
}

type TimeInformation struct {
	weekInfo *WeekInformation
}

func CreateTimeInformation() *TimeInformation {
	return &TimeInformation{
		weekInfo: &WeekInformation{},
	}
}

func (wi *WeekInformation) incrementWeekInformation() {
	wi.mutex.Lock()
	wi.weekTypeNumber = (wi.weekTypeNumber + 1) % weekTypesCount
	wi.weekTypeStr = WeekTypes[wi.weekTypeNumber]
	wi.mutex.Unlock()
}

func (timer *TimeInformation) UpdateWeekType() {
	go func(_timer *TimeInformation) {
		//added signal end of work
		for alive := true; alive; {
			// Sunday Weekday = iota
			// Monday 1
			// Tuesday 2
			// Wednesday 3
			// Thursday 4
			// Friday 5
			// Saturday 6
			getDaysToMonday()
			timer := time.NewTimer(5 * time.Second)
			select {
			case <-timer.C:
				_timer.weekInfo.incrementWeekInformation()
			}
		}
	}(timer)
}

func getDaysToMonday() int {
	for i := 0; i < 8; i++ {

		weekday := time.Now().Add(time.Duration(i) * 24 * time.Hour).Weekday()
		if weekday == time.Monday {
			return int(weekday)
		}
	}

	return -1
}

func (timer *TimeInformation) GetTodayDayNumber() (string, int) {
	data := time.Now().Weekday()
	return data.String(), int(data)
}

func (timer *TimeInformation) GetTomorrowDayNumberAndWeekType() (string, int, string, int) {
	dayData := time.Now().Add(24 * time.Hour).Weekday()
	timer.weekInfo.mutex.RLock()
	timer.weekInfo.mutex.RUnlock()
	return dayData.String(), int(dayData), timer.weekInfo.weekTypeStr, timer.weekInfo.weekTypeNumber
}

func (timer *TimeInformation) GetCurrentWeekType() *WeekInformation {
	timer.weekInfo.mutex.RLock()
	defer timer.weekInfo.mutex.RUnlock()
	return timer.weekInfo
}

func (timer *TimeInformation) IdentifyCurrentPair(timeClass map[int8]filesapi.TimeClasses) (string, error) {
	currentTime := time.Date(1, 1, 1, time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)

	if _, dayNumber := timer.GetTodayDayNumber(); dayNumber == 0 {
		return "Сегодня пар нет", nil
	}

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
			return fmt.Sprintf("Следующая пара %d в %.2d:%.2d:%.2d", i+1, hours, minutes, seconds), nil
		} else if currentTime.After(timeFrom) && currentTime.Before(timeTo) {
			hours, minutes, seconds := timeTo.Clock()
			return fmt.Sprintf("Сейчас идет пара %d до %.2d:%.2d:%.2d", i+1, hours, minutes, seconds), nil
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
