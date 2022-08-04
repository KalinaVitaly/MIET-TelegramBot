package timeapi

import (
	"fmt"
	"sync"
	"time"
)

const (
	weekTypesCount = 4
)

var weekTypes [weekTypesCount]string

func init() {
	weekTypes = [...]string{
		"1 числитель",
		"2 числитель",
		"1 знаменатель",
		"2 знаменатель",
	}
}

type WeekInformation struct {
	WeekTypeStr    string
	WeekTypeNumber int
	mutex          sync.RWMutex
}

func (wi *WeekInformation) incrementWeekInformation() {
	wi.mutex.Lock()
	wi.WeekTypeNumber = (wi.WeekTypeNumber + 1) % weekTypesCount
	wi.WeekTypeStr = weekTypes[wi.WeekTypeNumber]
	wi.mutex.Unlock()
}

func getTimeToMonday() (time.Duration, error) {
	var timeToMonday time.Duration
	for i := 0; i < 8; i++ {

		weekday := time.Now().Add(time.Duration(i) * 24 * time.Hour).Weekday()
		if weekday == time.Monday {
			timeToMonday =
				time.Until(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+i, 0, 0, 0, 0, time.Local))
			return timeToMonday, nil
		}
	}

	return timeToMonday, fmt.Errorf("Error calc weekday")
}

func (wi *WeekInformation) GetCurrentWeekType() *WeekInformation {
	wi.mutex.RLock()
	defer wi.mutex.RUnlock()
	return wi
}
