package datareader

import (
	"errors"
)

type ScheduleUniversity struct {
	GroupsSchedule map[string]*GroupSchedule
	ClassTime      map[int8]TimeClass
	CurrentWeek    int8
}

func CreateScheduleUniversity(dirPath string) (*ScheduleUniversity, error) {

	groupSchdule, err := ReadFilesFromDir(dirPath)

	if err != nil {
		return nil, err
	}

	classTime := make(map[int8]TimeClass, 7)
	groupScheduleMap := make(map[string]*GroupSchedule)

	if len(groupSchdule) <= 0 {
		return nil, errors.New("Error: read GroupSchedule failed!")
	}
	//schedule.Data[0].Group.Name

	//Get schedule time lessons
	for i := range groupSchdule[0].Times {
		classTime[int8(i)] = groupSchdule[0].Times[i]
	}

	for i := range groupSchdule {
		groupScheduleMap[groupSchdule[i].Data[0].Group.Name] = groupSchdule[i]
	}

	return &ScheduleUniversity{
		ClassTime:      classTime,
		GroupsSchedule: groupScheduleMap,
	}, nil
}
