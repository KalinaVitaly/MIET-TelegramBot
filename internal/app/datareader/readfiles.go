package datareader

import (
	"fmt"
	"io/ioutil"
)

func ReadFilesFromDir(dirPath string) ([]*GroupSchedule, error) {
	var groupScheduleList []*GroupSchedule
	filespathes, err := getFileNameFromDir(dirPath)
	if err != nil {
		fmt.Println("Error : read file names from dir failed!")
		return groupScheduleList, err
	}

	for _, value := range filespathes {
		schedule, err := ReadDataFromFile(dirPath + value)

		if err != nil {
			fmt.Println("Error : read data from file failed!")
			return groupScheduleList, err
		}

		groupScheduleList = append(groupScheduleList, schedule)
	}

	return groupScheduleList, nil
}

func ReadDataFromFile(filePath string) (*GroupSchedule, error) {
	fileData, err := ioutil.ReadFile(filePath)
	// can file be opened?
	if err != nil {
		return nil, err
	}

	result, err := NewGroupSchedule(fileData)

	if err != nil {
		return nil, err
	}

	return result, err
}

func getFileNameFromDir(dirPath string) ([]string, error) {
	var fileList []string
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fileList, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList, nil
}
