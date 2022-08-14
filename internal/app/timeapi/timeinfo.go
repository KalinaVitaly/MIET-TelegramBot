package timeapi

import (
	"MIET-TelegramBot/internal/app/filesapi"
	"fmt"
	"log"
	"time"
)

type TimeInformation struct {
	WeekInfo *WeekInformation
}

func CreateTimeInformation() *TimeInformation {
	return &TimeInformation{
		WeekInfo: &WeekInformation{
			WeekTypeNumber: 0,
			WeekTypeStr:    weekTypes[0],
		},
	}
}

func (timer *TimeInformation) UpdateWeekType() {
	go func(_timer *TimeInformation) {
		//added signal end of work
		for alive := true; alive; {
			timeToMonday, err := getTimeToMonday()

			if err != nil {
				log.Println(fmt.Sprintln("Error %s", err.Error))
			}
			timer := time.NewTimer(timeToMonday)
			select {
			case <-timer.C:
				_timer.WeekInfo.incrementWeekInformation()
			}
		}
	}(timer)
}

func (timer *TimeInformation) GetTodayDayNumber() (string, int) {
	data := time.Now().Weekday()
	return data.String(), (int(data) + 6) % 7
}

func (timer *TimeInformation) GetTomorrowDayNumberAndWeekType() (string, int, string, int) {
	dayData := time.Now().Add(24 * time.Hour).Weekday()

	tomorrowNumber := (int(dayData) + 6) % 7
	tomorrowString := dayData.String()

	timer.WeekInfo.mutex.RLock()

	weekTypeString := timer.WeekInfo.WeekTypeStr
	weekTypeNumber := timer.WeekInfo.WeekTypeNumber

	timer.WeekInfo.mutex.RUnlock()

	if tomorrowNumber == 1 {
		weekTypeNumber = (weekTypeNumber + 1) % weekTypesCount
		weekTypeString = weekTypes[weekTypeNumber]
	}

	return tomorrowString, tomorrowNumber, weekTypeString, weekTypeNumber
}

func (timer *TimeInformation) GetCurrentWeekType() *WeekInformation {
	timer.WeekInfo.mutex.RLock()
	defer timer.WeekInfo.mutex.RUnlock()
	return timer.WeekInfo
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
