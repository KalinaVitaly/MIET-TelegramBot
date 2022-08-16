package tools

import (
	"fmt"
	"log"
	"time"
)

func ConvertStringToTime(timeStr string) (time.Time, error) {
	timeT, err := time.Parse(time.RFC3339, timeStr+"Z")

	if err != nil {
		log.Println(fmt.Sprintf("Error convert string time to time type %s %s", err.Error(), timeStr))
		return timeT, err
	}

	return timeT, nil
}
