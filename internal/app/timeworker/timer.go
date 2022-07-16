package timeworker

import (
	"MIET-TelegramBot/internal/app/datareader"
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

func createWeekInformation() *WeekInformation {
	weekTypes := make(map[int]string)
	weekTypes[1] = FirstNumerator
}

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

func (time *Timer) IdentifyCurrentPair(timeClass map[int8]datareader.TimeClass) int {
	for i := 0; i < len(timeClass); i++ {
		if i == 0 && timeClass[i].TimeFrom
	}

}

func convertStringToTime(data string) (time.Time, error) {
	dataT, err := time.Parse("10:10:10", data)

	if err != nil {
		log.Println(fmt.Sprintf("Error convert string time to time type %s", err.Error()
		return nil, err
	}

	return dataT, nil
}
