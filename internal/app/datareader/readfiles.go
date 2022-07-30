package datareader

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func ReadFilesFromDir(dirPath string) ([]*GroupSchedule, error) {
	var groupScheduleList []*GroupSchedule
	chanReadFile := make(chan *fileTransferData)
	filesPathes, err := getFileNameFromDir(dirPath)

	if err != nil {
		fmt.Println("Error : read file names from dir failed!")
		return groupScheduleList, err
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover after panic ! ", err)
		}

		close(chanReadFile)
	}()

	for _, value := range filesPathes {
		go readDataFromFile(value, chanReadFile)
	}

	var errorCounter int

	for {
		select {
		case <-time.After(time.Second * 10):
			log.Println("Something went wrong...")
			return nil, fmt.Errorf("Long reading file")
		case data := <-chanReadFile:
			if data.containsError() {
				log.Println(fmt.Sprintf("File %s read failed. Error: %s", data.filePath, data.readError.Error()))
				errorCounter++
			} else {
				log.Println(fmt.Sprintf("File %s successfully read ", data.filePath))
				groupScheduleList = append(groupScheduleList, data.groupsSchedule)
			}

			log.Println(errorCounter, len(groupScheduleList), len(filesPathes), errorCounter+len(groupScheduleList))
			if errorCounter+len(groupScheduleList) == len(filesPathes) {
				log.Println(fmt.Sprintf("File successfully read %d and failed read %d ", len(groupScheduleList), errorCounter))

				if errorCounter != 0 {
					return groupScheduleList, fmt.Errorf("Error read %d file/files", errorCounter)
				}
				return groupScheduleList, nil
			}
		}
	}
}

func readDataFromFile(filePath string, chanReadFile chan<- *fileTransferData) {
	fileData, err := ioutil.ReadFile(filePath)
	transferData := newFileTransferData(nil, filePath, nil)
	if err != nil {
		transferData.readError = err
		chanReadFile <- transferData
		return
	}

	result, err := NewGroupSchedule(fileData)

	if err != nil {
		transferData.readError = err
		chanReadFile <- transferData
		return
	}

	transferData.groupsSchedule = result
	chanReadFile <- transferData
	return
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
