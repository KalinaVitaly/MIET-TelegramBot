package datareader

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFilesFromDir(dirPath string) ([]*GroupSchedule, error) {
	var groupScheduleList []*GroupSchedule
	errorChan := make(chan error)
	resultChan := make(chan *GroupSchedule)
	filespathes, err := getFileNameFromDir(dirPath)

	if err != nil {
		fmt.Println("Error : read file names from dir failed!")
		return groupScheduleList, err
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover after panic ! ", err)
		}
	}()

	for _, value := range filespathes {
		go readDataFromFile(value, errorChan, resultChan)
	}

	for {
		select {
		case err := <-errorChan:
			log.Println("Error read file :", err)
		case data := <-resultChan:
			groupScheduleList = append(groupScheduleList, data)

			if len(groupScheduleList) == len(filespathes) {
				log.Println("All files read ", len(groupScheduleList), len(filespathes))
				return groupScheduleList, nil
			}
		}
	}
}

func readDataFromFile(filePath string, chError chan<- error, chResult chan<- *GroupSchedule) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		chError <- err
	}

	result, err := NewGroupSchedule(fileData)

	if err != nil {
		chError <- err
	}

	chResult <- result
}

func getFileNameFromDir(dirPath string) ([]string, error) {
	var fileList []string
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return fileList, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, dirPath+file.Name())
		}
	}

	return fileList, nil
}
