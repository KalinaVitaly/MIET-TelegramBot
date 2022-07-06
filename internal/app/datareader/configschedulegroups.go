package datareader

import (
	"errors"
	"strings"
)

type ScheduleUniversity struct {
	GroupsSchedule map[string]*GroupSchedule
	ClassTime      map[int8]TimeClass
	CurrentWeek    int8
}

var scheduleUniversity *ScheduleUniversity

func CreateScheduleUniversity(dirPath string) error {

	groupSchdule, err := ReadFilesFromDir(dirPath)

	if err != nil {
		return err
	}

	classTime := make(map[int8]TimeClass, 7)
	groupScheduleMap := make(map[string]*GroupSchedule)

	if len(groupSchdule) <= 0 {
		return errors.New("Error: read GroupSchedule failed!")
	}

	for i := range groupSchdule[0].Times {
		classTime[int8(i)] = groupSchdule[0].Times[i]
	}

	for i := range groupSchdule {
		groupScheduleMap[groupSchdule[i].Data[0].Group.Name] = groupSchdule[i]
	}

	scheduleUniversity = &ScheduleUniversity{
		ClassTime:      classTime,
		GroupsSchedule: groupScheduleMap,
	}
	return nil
}

func (scheduleGroups *ScheduleUniversity) ClassTimeToString() (result string) {
	for i := 0; i < len(scheduleGroups.ClassTime); i++ {
		result += scheduleGroups.ClassTime[int8(i)].Time + " Начало : " + scheduleGroups.ClassTime[int8(i)].TimeFrom + " Конец : " + scheduleGroups.ClassTime[int8(i)].TimeTo + "\n"
	}

	result = strings.Replace(result, "0001-01-01T", "", -1)

	return result
}

func GetInstanceScheduleUniversity() *ScheduleUniversity {
	return scheduleUniversity
}
