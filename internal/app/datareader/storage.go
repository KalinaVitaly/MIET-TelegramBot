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

func NewGroupSchedule(data []byte) (*GroupSchedule, error) {
	var groupSchedule GroupSchedule

	if err := json.Unmarshal(data, &groupSchedule); err != nil {
		fmt.Println("Error: Unmarshal group schedule failed!")
		return nil, err
	}

	return &groupSchedule, nil
}
