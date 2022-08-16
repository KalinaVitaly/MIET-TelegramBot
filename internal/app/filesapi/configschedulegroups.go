package filesapi

import (
	"MIET-TelegramBot/internal/app/tools"
	"fmt"
	"log"
	"strings"
)

type ScheduleUniversity struct {
	GroupsSchedule map[string]*GroupSchedule
	ClassTime      map[int8]TimeClasses
	CurrentWeek    int8
}

func CreateScheduleUniversity(dirPath string) (*ScheduleUniversity, error) {
	groupSchedule, err := readFilesFromDir(dirPath)
	const classesCount = 7

	if err != nil {
		return nil, err
	}

	classTime := make(map[int8]TimeClasses, classesCount)
	groupScheduleMap := make(map[string]*GroupSchedule)

	for i := range groupSchedule[0].Times {
		classTime[int8(i)] = groupSchedule[0].Times[i]
	}

	for i := range groupSchedule {
		log.Println(groupSchedule[i].Data[0].Group.Name)
		groupScheduleMap[groupSchedule[i].Data[0].Group.Name] = groupSchedule[i]
	}

	return &ScheduleUniversity{
			ClassTime:      classTime,
			GroupsSchedule: groupScheduleMap,
		},
		nil
}

func (s *ScheduleUniversity) GetClassesCurrentWeek(group string, weekType int) string {
	var weekSchedule string

	for i := 0; i < 7; i++ {
		weekSchedule += s.GetClassesInSelectedDay(group, i, weekType) + "\n"
	}

	return weekSchedule
}

func (s *ScheduleUniversity) GetShortClassesCurrentWeek(group string, weekType int) (string, error) {
	var weekSchedule string

	for i := 0; i < 7; i++ {
		dailySchedule, err := s.getShortClassesInSelectedDay(group, i, weekType)
		if err != nil {
			return "", err
		}
		weekSchedule += fmt.Sprintln(dailySchedule)
	}

	return weekSchedule, nil
}

func (s *ScheduleUniversity) getShortClassesInSelectedDay(group string, dayNumber, weekType int) (string, error) {
	var todaySchedule string

	for _, value := range s.GroupsSchedule[group].Data {
		if value.Day == dayNumber && value.DayNumber == weekType {

			timeFrom, err := tools.ConvertStringToTime(value.Time.TimeFrom)

			if err != nil {
				log.Println("Error convert string to time")
				return "", err
			}

			timeTo, err := tools.ConvertStringToTime(value.Time.TimeTo)

			if err != nil {
				log.Println("Error convert string to time")
				return "", err
			}

			hoursFrom, minutesFrom, secondsFrom := timeFrom.Clock()
			hoursTo, minutesTo, secondsTo := timeTo.Clock()

			todaySchedule += fmt.Sprintln(fmt.Sprintf("%.2d:%.2d:%.2d  %.2d:%.2d:%.2d", hoursFrom, minutesFrom, secondsFrom, hoursTo, minutesTo, secondsTo))
			todaySchedule += fmt.Sprintln(value.Class.Name)
			todaySchedule += fmt.Sprintln("Кабинет : " + value.Room.Name)

			continue
		}
	}
	return todaySchedule, nil
}

func (s *ScheduleUniversity) GetClassesInSelectedDay(group string, dayNumber, weekType int) string {
	var todaySchedule string

	for _, value := range s.GroupsSchedule[group].Data {
		if value.Day == dayNumber && value.DayNumber == weekType {
			todaySchedule += value.Time.TimeFrom + "\n"
			todaySchedule += value.Class.Name + "\n"
			todaySchedule += value.Class.TeacherFull + "\n"
			todaySchedule += value.Class.Form + "\n"
			todaySchedule += "Кабинет : " + value.Room.Name + "\n\n"

			continue
		}
	}
	return todaySchedule
}

func (scheduleGroups *ScheduleUniversity) ClassTimeToString() (result string) {
	for i := 0; i < len(scheduleGroups.ClassTime); i++ {
		result += scheduleGroups.ClassTime[int8(i)].Time + " Начало : " + scheduleGroups.ClassTime[int8(i)].TimeFrom + " Конец : " + scheduleGroups.ClassTime[int8(i)].TimeTo + "\n"
	}

	result = strings.Replace(result, "0001-01-01T", "", -1)

	return result
}

func (scheduleGroups *ScheduleUniversity) GetTimeClass() map[int8]TimeClasses {
	return scheduleGroups.ClassTime
}
