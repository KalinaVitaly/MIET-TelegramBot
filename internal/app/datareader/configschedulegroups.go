package datareader

import "errors"

type ScheduleUniversity struct {
	GroupsSchedule map[string]*GroupSchedule
	ClassTime      map[int8]TimeClass
	CurrentWeek    int8
}

func CreateScheduleUniversity(data []*GroupSchedule) (*ScheduleUniversity, error) {
	classTime := make(map[int8]TimeClass, 7)
	groupSchedule := make(map[string]*GroupSchedule)

	if len(data) <= 0 {
		return nil, errors.New("Error: read GroupSchedule failed!")
	}
	//schedule.Data[0].Group.Name

	//Get schedule time lessons
	for i := range data[0].Times {
		classTime[int8(i)] = data[0].Times[i]
	}

	for i := range data {
		groupSchedule[data[i].Data[0].Group.Name] = data[i]
	}

	return &ScheduleUniversity{
		ClassTime:      classTime,
		GroupsSchedule: groupSchedule,
	}, nil
}
