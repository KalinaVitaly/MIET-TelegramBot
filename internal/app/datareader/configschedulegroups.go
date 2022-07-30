package datareader

import "strings"

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
		groupScheduleMap[groupSchedule[i].Data[0].Group.Name] = groupSchedule[i]
	}

	return &ScheduleUniversity{
			ClassTime:      classTime,
			GroupsSchedule: groupScheduleMap,
		},
		nil
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
